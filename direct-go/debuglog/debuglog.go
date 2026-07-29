// Package debuglog provides a logger that buffers metadata-safe diagnostics
// and optionally forwards them to a debug server.
package debuglog

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Log levels.
const (
	LevelOff     = 0 // No debug logging
	LevelNormal  = 1 // Important debug messages
	LevelVerbose = 2 // All debug messages including ping/pong
)

const (
	defaultLogBufferSize  = 5000
	defaultLogChannelSize = 100
)

// LogEntry represents a structured log message.
type LogEntry struct {
	Time    time.Time `json:"time"`
	Level   string    `json:"level"`
	Message string    `json:"message"`
	Source  string    `json:"source,omitempty"`
}

// LogQuery defines filters for querying logs.
type LogQuery struct {
	Level   string
	Keyword string
	Limit   int
	Since   time.Time
}

// RingBuffer holds a fixed number of log entries.
type RingBuffer struct {
	entries []LogEntry
	head    int
	size    int
	cap     int
	mu      sync.RWMutex
}

func NewRingBuffer(capacity int) *RingBuffer {
	if capacity <= 0 {
		capacity = 1
	}
	return &RingBuffer{
		entries: make([]LogEntry, capacity),
		cap:     capacity,
	}
}

// Add adds a new entry to the ring buffer.
func (rb *RingBuffer) Add(entry LogEntry) {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	rb.entries[rb.head] = entry
	rb.head = (rb.head + 1) % rb.cap
	if rb.size < rb.cap {
		rb.size++
	}
}

// Clear removes all entries from the ring buffer.
func (rb *RingBuffer) Clear() {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	rb.entries = make([]LogEntry, rb.cap)
	rb.head = 0
	rb.size = 0
}

// Query returns logs matching the query.
func (rb *RingBuffer) Query(q LogQuery) []LogEntry {
	rb.mu.RLock()
	defer rb.mu.RUnlock()

	var results []LogEntry
	start := (rb.head - rb.size + rb.cap) % rb.cap
	for i := 0; i < rb.size; i++ {
		entry := rb.entries[(start+i)%rb.cap]
		if !q.Since.IsZero() && entry.Time.Before(q.Since) {
			continue
		}
		if q.Level != "" && !strings.EqualFold(entry.Level, q.Level) {
			continue
		}
		if q.Keyword != "" && !strings.Contains(strings.ToLower(entry.Message), strings.ToLower(q.Keyword)) {
			continue
		}
		results = append(results, entry)
	}
	if q.Limit > 0 && len(results) > q.Limit {
		results = results[len(results)-q.Limit:]
	}
	return results
}

// LoggerOptions configures an independent debug logger.
//
// Safe logging is the default. Payload-level tracing must be enabled explicitly
// with EnableUnsafePayloadTracing and should only be used in an isolated,
// controlled environment.
type LoggerOptions struct {
	Level                int
	ServerURL            string
	UnsafePayloadTracing bool
	Writer               io.Writer
}

// Logger buffers metadata-safe diagnostics and optionally forwards them to a
// debug server. Logger instances are independent, so applications can scope
// debug configuration to a single client instead of changing process-global
// state.
type Logger struct {
	buffer      *RingBuffer
	subscribers map[chan LogEntry]struct{}
	serverURL   string
	enabled     bool
	level       int
	unsafe      bool
	localLogger *log.Logger
	mu          sync.RWMutex
	subMu       sync.RWMutex
}

// NewLogger creates an independent debug logger.
func NewLogger(opts LoggerOptions) *Logger {
	writer := opts.Writer
	if writer == nil {
		writer = os.Stdout
	}
	level := opts.Level
	if level < LevelOff {
		level = LevelOff
	}
	if level > LevelVerbose {
		level = LevelVerbose
	}
	serverURL := strings.TrimRight(strings.TrimSpace(opts.ServerURL), "/")
	if serverURL != "" && level == LevelOff {
		// Supplying a server URL is an explicit request to collect the safe
		// diagnostics. Payload tracing remains disabled independently.
		level = LevelNormal
	}
	l := &Logger{
		buffer:      NewRingBuffer(defaultLogBufferSize),
		subscribers: make(map[chan LogEntry]struct{}),
		serverURL:   serverURL,
		enabled:     serverURL != "",
		level:       level,
		unsafe:      opts.UnsafePayloadTracing,
		localLogger: log.New(writer, "", log.LstdFlags),
	}
	if l.unsafe {
		l.localLogger.Print("[WARNING] UNSAFE payload tracing enabled; debug logs may contain credentials and private message data")
	}
	return l
}

func defaultLoggerOptions() LoggerOptions {
	opts := LoggerOptions{}
	if v := os.Getenv("DIRECT_DEBUG"); v != "" {
		if level, err := strconv.Atoi(v); err == nil {
			opts.Level = level
		} else if strings.EqualFold(v, "true") {
			opts.Level = LevelNormal
		}
	}
	opts.UnsafePayloadTracing = strings.EqualFold(os.Getenv("DIRECT_DEBUG_UNSAFE_PAYLOADS"), "true")
	return opts
}

var defaultLogger = NewLogger(defaultLoggerOptions())

// Default returns the process-wide logger used by the legacy package-level API.
// New clients use this logger unless Options.DebugLogger is provided.
func Default() *Logger {
	return defaultLogger
}

// Subscribe adds a channel to receive real-time logs.
func (l *Logger) Subscribe() chan LogEntry {
	ch := make(chan LogEntry, defaultLogChannelSize)
	l.subMu.Lock()
	l.subscribers[ch] = struct{}{}
	l.subMu.Unlock()
	return ch
}

// Unsubscribe removes a subscriber channel.
func (l *Logger) Unsubscribe(ch chan LogEntry) {
	l.subMu.Lock()
	delete(l.subscribers, ch)
	close(ch)
	l.subMu.Unlock()
}

// GetBuffer returns the logger's internal ring buffer.
func (l *Logger) GetBuffer() *RingBuffer {
	return l.buffer
}

// Broadcast sends an entry to all subscribers.
func (l *Logger) Broadcast(entry LogEntry) {
	l.subMu.RLock()
	defer l.subMu.RUnlock()
	for ch := range l.subscribers {
		select {
		case ch <- entry:
		default:
		}
	}
}

// GetLogs returns logs matching the query.
func (l *Logger) GetLogs(q LogQuery) []LogEntry {
	return l.buffer.Query(q)
}

// SetServer sets the debug server URL and enables remote logging.
func (l *Logger) SetServer(url string) {
	url = strings.TrimRight(strings.TrimSpace(url), "/")
	l.mu.Lock()
	l.serverURL = url
	l.enabled = url != ""
	if l.enabled && l.level == LevelOff {
		// Keep EnableDebugServer useful as an explicit, safe opt-in while
		// preserving LevelOff when no server was requested.
		l.level = LevelNormal
	}
	l.mu.Unlock()
}

// SetLevel changes the logger's debug log level. LevelNormal and LevelVerbose
// only emit metadata-safe diagnostics from the SDK.
func (l *Logger) SetLevel(level int) {
	if level < LevelOff {
		level = LevelOff
	}
	if level > LevelVerbose {
		level = LevelVerbose
	}
	l.mu.Lock()
	l.level = level
	l.mu.Unlock()
}

// Level returns the current logger's debug log level.
func (l *Logger) Level() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.level
}

// EnableUnsafePayloadTracing enables logs that may contain message bodies or
// complete protocol payloads. It is intended only for isolated, controlled
// debugging and is disabled by default.
func (l *Logger) EnableUnsafePayloadTracing() {
	l.mu.Lock()
	l.unsafe = true
	l.mu.Unlock()
	l.localLogger.Print("[WARNING] UNSAFE payload tracing enabled; debug logs may contain credentials and private message data")
}

// DisableUnsafePayloadTracing disables unsafe payload tracing.
func (l *Logger) DisableUnsafePayloadTracing() {
	l.mu.Lock()
	l.unsafe = false
	l.mu.Unlock()
}

// UnsafePayloadTracingEnabled reports whether unsafe payload tracing is enabled.
func (l *Logger) UnsafePayloadTracingEnabled() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.unsafe
}

// Clear removes all buffered debug entries.
func (l *Logger) Clear() {
	l.buffer.Clear()
}

// Printf logs a metadata-safe message (level 1 = normal). Callers must pass
// only metadata-safe values. Use UnsafePrintf for intentional payload tracing.
func (l *Logger) Printf(format string, v ...interface{}) {
	l.logMessage(LevelNormal, "info", format, v...)
}

// Verbose logs a metadata-safe message at verbose level (level 2).
func (l *Logger) Verbose(format string, v ...interface{}) {
	l.logMessage(LevelVerbose, "debug", format, v...)
}

// UnsafePrintf emits a payload-level normal diagnostic only after explicit
// unsafe tracing has been enabled with EnableUnsafePayloadTracing.
func (l *Logger) UnsafePrintf(format string, v ...interface{}) {
	if !l.UnsafePayloadTracingEnabled() {
		return
	}
	l.logMessage(LevelNormal, "info", format, v...)
}

// UnsafeVerbose emits a payload-level verbose diagnostic only after explicit
// unsafe tracing has been enabled with EnableUnsafePayloadTracing.
func (l *Logger) UnsafeVerbose(format string, v ...interface{}) {
	if !l.UnsafePayloadTracingEnabled() {
		return
	}
	l.logMessage(LevelVerbose, "debug", format, v...)
}

var httpClient = &http.Client{Timeout: 1 * time.Second}

func (l *Logger) logMessage(level int, levelStr string, format string, v ...interface{}) {
	l.mu.RLock()
	currentLevel := l.level
	url := l.serverURL
	on := l.enabled
	l.mu.RUnlock()
	if currentLevel < level {
		return
	}

	msg := fmt.Sprintf(format, v...)
	l.localLogger.Print(msg)
	entry := LogEntry{Time: time.Now(), Level: levelStr, Message: msg}
	l.buffer.Add(entry)
	l.Broadcast(entry)

	if on && url != "" {
		go func() {
			data, err := json.Marshal(entry)
			if err != nil {
				return
			}
			resp, err := httpClient.Post(url+"/log", "application/json", bytes.NewBuffer(data))
			if err == nil {
				resp.Body.Close()
			}
		}()
	}
}

// RedactID returns a stable, non-reversible identifier suitable for correlating
// logs without exposing the original user, talk, or message ID.
func RedactID(value interface{}) string {
	raw := strings.TrimSpace(fmt.Sprint(value))
	if raw == "" {
		return "<empty>"
	}
	hash := sha256.Sum256([]byte(raw))
	return "id:" + hex.EncodeToString(hash[:])[:12]
}

// RedactSecret returns a placeholder without returning a secret or secret
// reference. It intentionally does not include length, hash, or path data.
func RedactSecret(value string) string {
	if strings.TrimSpace(value) == "" {
		return "<empty secret>"
	}
	return "<redacted secret>"
}

// RedactAuthorization removes an authorization header value from diagnostics.
func RedactAuthorization(value string) string {
	if strings.TrimSpace(value) == "" {
		return "<empty authorization>"
	}
	return "<redacted authorization>"
}

// SummarizePayload returns metadata about a structured value without including
// strings, map values, or other payload contents.
func SummarizePayload(value interface{}) string {
	if value == nil {
		return "nil"
	}
	return summarizeValue(reflect.ValueOf(value))
}

func summarizeValue(value reflect.Value) string {
	if !value.IsValid() {
		return "nil"
	}
	for value.Kind() == reflect.Interface || value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return "nil"
		}
		value = value.Elem()
	}
	switch value.Kind() {
	case reflect.String:
		return fmt.Sprintf("string(len=%d)", value.Len())
	case reflect.Slice, reflect.Array:
		return fmt.Sprintf("%s(len=%d)", value.Kind(), value.Len())
	case reflect.Map:
		return fmt.Sprintf("map(len=%d)", value.Len())
	case reflect.Struct:
		return fmt.Sprintf("struct(%s)", value.Type().String())
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "integer"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return "unsigned"
	case reflect.Float32, reflect.Float64:
		return "float"
	default:
		return value.Kind().String()
	}
}

// Println logs a metadata-safe message with a newline.
func (l *Logger) Println(v ...interface{}) {
	l.Printf("%s", fmt.Sprintln(v...))
}

// Writer returns an io.Writer that records only the number of bytes written.
// It is safe to use with arbitrary streams because it never records their
// contents. Use UnsafeWriter only for isolated troubleshooting after enabling
// unsafe payload tracing.
func (l *Logger) Writer() io.Writer {
	return &debugWriter{logger: l}
}

// UnsafeWriter returns an io.Writer that forwards raw bytes only when unsafe
// payload tracing has been explicitly enabled.
func (l *Logger) UnsafeWriter() io.Writer {
	return &unsafeDebugWriter{logger: l}
}

type debugWriter struct{ logger *Logger }

func (w *debugWriter) Write(p []byte) (n int, err error) {
	w.logger.Printf("writer bytes=%d", len(p))
	return len(p), nil
}

type unsafeDebugWriter struct{ logger *Logger }

func (w *unsafeDebugWriter) Write(p []byte) (n int, err error) {
	w.logger.UnsafePrintf("%s", string(p))
	return len(p), nil
}

// Package-level functions preserve the original API and use Default.
func Subscribe() chan LogEntry                      { return defaultLogger.Subscribe() }
func Unsubscribe(ch chan LogEntry)                  { defaultLogger.Unsubscribe(ch) }
func GetBuffer() *RingBuffer                        { return defaultLogger.GetBuffer() }
func Broadcast(entry LogEntry)                      { defaultLogger.Broadcast(entry) }
func GetLogs(q LogQuery) []LogEntry                 { return defaultLogger.GetLogs(q) }
func SetServer(url string)                          { defaultLogger.SetServer(url) }
func SetLevel(level int)                            { defaultLogger.SetLevel(level) }
func Level() int                                    { return defaultLogger.Level() }
func EnableUnsafePayloadTracing()                   { defaultLogger.EnableUnsafePayloadTracing() }
func DisableUnsafePayloadTracing()                  { defaultLogger.DisableUnsafePayloadTracing() }
func UnsafePayloadTracingEnabled() bool             { return defaultLogger.UnsafePayloadTracingEnabled() }
func Clear()                                        { defaultLogger.Clear() }
func Printf(format string, v ...interface{})        { defaultLogger.Printf(format, v...) }
func Verbose(format string, v ...interface{})       { defaultLogger.Verbose(format, v...) }
func UnsafePrintf(format string, v ...interface{})  { defaultLogger.UnsafePrintf(format, v...) }
func UnsafeVerbose(format string, v ...interface{}) { defaultLogger.UnsafeVerbose(format, v...) }
func Println(v ...interface{})                      { defaultLogger.Println(v...) }
func Writer() io.Writer                             { return defaultLogger.Writer() }
func UnsafeWriter() io.Writer                       { return defaultLogger.UnsafeWriter() }
