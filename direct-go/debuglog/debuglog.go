// Package debuglog provides a logger that sends logs to a debug server
package debuglog

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

// Log levels
const (
	LevelOff     = 0 // No debug logging
	LevelNormal  = 1 // Important debug messages
	LevelVerbose = 2 // All debug messages including ping/pong
)

var (
	debugServerURL string
	enabled        bool
	logLevel       int
	client         = &http.Client{Timeout: 1 * time.Second}
	mu             sync.Mutex
	localLogger    = log.New(os.Stdout, "", log.LstdFlags)
)

func init() {
	// Check DIRECT_DEBUG environment variable
	// 0 or unset: no logging
	// 1: normal logging
	// 2: verbose logging (includes ping/pong)
	if v := os.Getenv("DIRECT_DEBUG"); v != "" {
		if level, err := strconv.Atoi(v); err == nil {
			logLevel = level
		} else if v == "true" {
			logLevel = LevelNormal
		}
	}
}

// SetServer sets the debug server URL and enables remote logging
func SetServer(url string) {
	mu.Lock()
	defer mu.Unlock()
	debugServerURL = url
	enabled = url != ""
}

// IsEnabled returns true if remote logging is enabled
func IsEnabled() bool {
	mu.Lock()
	defer mu.Unlock()
	return enabled
}

// Printf logs a message (level 1 = normal)
func Printf(format string, v ...interface{}) {
	printfLevel(LevelNormal, format, v...)
}

// Verbose logs a message at verbose level (level 2)
func Verbose(format string, v ...interface{}) {
	printfLevel(LevelVerbose, format, v...)
}

func printfLevel(level int, format string, v ...interface{}) {
	mu.Lock()
	currentLevel := logLevel
	url := debugServerURL
	on := enabled
	mu.Unlock()

	if currentLevel < level {
		return // Skip if log level is too low
	}

	msg := fmt.Sprintf(format, v...)
	localLogger.Print(msg)

	// Send to debug server if enabled
	if on && url != "" {
		go func() {
			resp, err := client.Post(url+"/log", "text/plain", bytes.NewBufferString(msg))
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

// Writer returns an io.Writer that sends output to both local and debug server
func Writer() io.Writer {
	return &debugWriter{}
}

type debugWriter struct{}

func (w *debugWriter) Write(p []byte) (n int, err error) {
	Printf("%s", string(p))
	return len(p), nil
}

