// Package debuglog provides a logger that sends logs to a debug server
package debuglog

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	debugServerURL string
	enabled        bool
	client         = &http.Client{Timeout: 1 * time.Second}
	mu             sync.Mutex
	localLogger    = log.New(os.Stdout, "", log.LstdFlags)
)

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

// Printf logs a message both locally and to the debug server
func Printf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)

	// Always log locally
	localLogger.Print(msg)

	// Send to debug server if enabled
	mu.Lock()
	url := debugServerURL
	on := enabled
	mu.Unlock()

	if on && url != "" {
		go func() {
			resp, err := client.Post(url+"/log", "text/plain", bytes.NewBufferString(msg))
			if err == nil {
				resp.Body.Close()
			}
		}()
	}
}

// Println logs a message with newline
func Println(v ...interface{}) {
	msg := fmt.Sprintln(v...)

	localLogger.Print(msg)

	mu.Lock()
	url := debugServerURL
	on := enabled
	mu.Unlock()

	if on && url != "" {
		go func() {
			resp, err := client.Post(url+"/log", "text/plain", bytes.NewBufferString(msg))
			if err == nil {
				resp.Body.Close()
			}
		}()
	}
}

// Writer returns an io.Writer that sends output to both local and debug server
func Writer() io.Writer {
	return &debugWriter{}
}

type debugWriter struct{}

func (w *debugWriter) Write(p []byte) (n int, err error) {
	msg := string(p)

	// Local output
	os.Stdout.Write(p)

	// Remote output
	mu.Lock()
	url := debugServerURL
	on := enabled
	mu.Unlock()

	if on && url != "" {
		go func() {
			resp, err := client.Post(url+"/log", "text/plain", bytes.NewBufferString(msg))
			if err == nil {
				resp.Body.Close()
			}
		}()
	}

	return len(p), nil
}
