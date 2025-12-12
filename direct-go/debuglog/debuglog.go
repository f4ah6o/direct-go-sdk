// Package debuglog provides a logger that buffers logs for the debug server
package debuglog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Log levels
const (
	LevelOff     = 0 // No debug logging
	LevelNormal  = 1 // Important debug messages
	LevelVerbose = 2 // All debug messages including ping/pong
)

// LogEntry represents a structured log message
type LogEntry struct {
	Time    time.Time `json:"time"`
	Level   string    `json:"level"` // "debug", "verbose", "info", "error"
	Message string    `json:"message"`
	Source  string    `json:"source,omitempty"`
}

// LogQuery defines filters for querying logs
type LogQuery struct {
	Level   string    // specific level to filter by
	Keyword string    // keyword to search in message
	Limit   int       // max number of entries
	Since   time.Time // return entries after this time
}

// RingBuffer holds a fixed number of log entries
type RingBuffer struct {
	entries []LogEntry
	head    int
	size    int
	cap     int
	mu      sync.RWMutex
}

func NewRingBuffer(capacity int) *RingBuffer {
	return &RingBuffer{
		entries: make([]LogEntry, capacity),
		cap:     capacity,
	}
}

// Add adds a new entry to the ring buffer
func (rb *RingBuffer) Add(entry LogEntry) {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	rb.entries[rb.head] = entry
	rb.head = (rb.head + 1) % rb.cap
	if rb.size < rb.cap {
		rb.size++
	}
}

// Query returns logs matching the query
func (rb *RingBuffer) Query(q LogQuery) []LogEntry {
	rb.mu.RLock()
	defer rb.mu.RUnlock()

	var results []LogEntry

	// Iterate through all valid entries
	// The oldest entry is at (head - size + cap) % cap
	start := (rb.head - rb.size + rb.cap) % rb.cap

	for i := 0; i < rb.size; i++ {
		idx := (start + i) % rb.cap
		entry := rb.entries[idx]

		// Apply filters
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

	// Apply limit (take last N)
	if q.Limit > 0 && len(results) > q.Limit {
		results = results[len(results)-q.Limit:]
	}

	return results
}

var (
	// buffer holds the latest logs
	buffer = NewRingBuffer(5000)

	// subscribers for real-time streaming
	subscribers = make(map[chan LogEntry]struct{})
	subMu       sync.RWMutex

	// debug server configuration
	debugServerURL string
	enabled        bool
	mu             sync.Mutex

	logLevel    int
	localLogger = log.New(os.Stdout, "", log.LstdFlags)
)

func init() {
	// Check DIRECT_DEBUG environment variable
	if v := os.Getenv("DIRECT_DEBUG"); v != "" {
		if level, err := strconv.Atoi(v); err == nil {
			logLevel = level
		} else if v == "true" {
			logLevel = LevelNormal
		}
	}
	enabled = true // Always enabled internally, just controls level
}

// Subscribe adds a channel to receive real-time logs
func Subscribe() chan LogEntry {
	ch := make(chan LogEntry, 100)
	subMu.Lock()
	subscribers[ch] = struct{}{}
	subMu.Unlock()
	return ch
}

// Unsubscribe removes a subscriber channel
func Unsubscribe(ch chan LogEntry) {
	subMu.Lock()
	delete(subscribers, ch)
	close(ch)
	subMu.Unlock()
}

// GetBuffer returns the internal ring buffer
func GetBuffer() *RingBuffer {
	return buffer
}

// Broadcast sends an entry to all subscribers
func Broadcast(entry LogEntry) {
	subMu.RLock()
	defer subMu.RUnlock()
	for ch := range subscribers {
		select {
		case ch <- entry:
		default:
			// fast non-blocking drop
		}
	}
}

// GetLogs returns logs matching the query
func GetLogs(q LogQuery) []LogEntry {
	return buffer.Query(q)
}

// SetServer sets the debug server URL and enables remote logging
func SetServer(url string) {
	mu.Lock()
	defer mu.Unlock()
	debugServerURL = url
	enabled = url != ""
}

// Printf logs a message (level 1 = normal)
func Printf(format string, v ...interface{}) {
	logMessage(LevelNormal, "info", format, v...)
}

// Verbose logs a message at verbose level (level 2)
func Verbose(format string, v ...interface{}) {
	logMessage(LevelVerbose, "debug", format, v...)
}

var httpClient = &http.Client{Timeout: 1 * time.Second}

func logMessage(level int, levelStr string, format string, v ...interface{}) {
	if logLevel < level {
		return
	}

	msg := fmt.Sprintf(format, v...)

	// Log to local stdout
	localLogger.Print(msg)

	// Create entry
	entry := LogEntry{
		Time:    time.Now(),
		Level:   levelStr,
		Message: msg,
	}

	// Add to local buffer
	buffer.Add(entry)

	// Broadcast to local subscribers
	subMu.RLock()
	for ch := range subscribers {
		select {
		case ch <- entry:
		default:
			// fast non-blocking drop
		}
	}
	subMu.RUnlock()

	// Send to remote server if enabled
	mu.Lock()
	url := debugServerURL
	on := enabled
	mu.Unlock()

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

// Println logs a message with newline (level 1 = normal)
func Println(v ...interface{}) {
	Printf("%s", fmt.Sprintln(v...))
}

// Writer returns an io.Writer that sends output to the log system
func Writer() io.Writer {
	return &debugWriter{}
}

type debugWriter struct{}

func (w *debugWriter) Write(p []byte) (n int, err error) {
	Printf("%s", string(p))
	return len(p), nil
}
