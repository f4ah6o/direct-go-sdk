package logserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/f4ah6o/direct-go-sdk/direct-go/debuglog"
)

// Server represents the log server
type Server struct {
	mux *http.ServeMux
}

// New creates a new log server
func New() *Server {
	s := &Server{
		mux: http.NewServeMux(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("/logs", s.handleLogs)
	s.mux.HandleFunc("/stream", s.handleStream)
	s.mux.HandleFunc("/", s.handleIndex)

	// Endpoint for collecting logs from other processes
	s.mux.HandleFunc("/log", s.handleLogPost)
}

// handleLogPost receives logs from other processes via HTTP POST
func (s *Server) handleLogPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var entry debuglog.LogEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		// Fallback for legacy plain text logs if needed, but we prefer JSON now
		// For now, strict JSON
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Override time if zero (trusted client time or server receipt time?)
	// Client time is better for correlation, but let's ensure it exists
	if entry.Time.IsZero() {
		entry.Time = time.Now()
	}

	// Add to our buffer
	debuglog.GetBuffer().Add(entry)

	// Broadcast to our subscribers
	debuglog.Broadcast(entry)

	w.WriteHeader(http.StatusOK)
}

// ListenAndServe starts the server
func (s *Server) ListenAndServe(addr string) error {
	fmt.Printf("Log server listening on %s\n", addr)
	return http.ListenAndServe(addr, s.mux)
}

// handleLogs returns logs as JSON
func (s *Server) handleLogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := debuglog.LogQuery{
		Level:   r.URL.Query().Get("level"),
		Keyword: r.URL.Query().Get("keyword"),
	}

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			query.Limit = limit
		}
	} else {
		query.Limit = 100 // Default limit
	}

	if sinceStr := r.URL.Query().Get("since"); sinceStr != "" {
		if t, err := time.Parse(time.RFC3339, sinceStr); err == nil {
			query.Since = t
		}
	}

	logs := debuglog.GetLogs(query)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"count":   len(logs),
		"entries": logs,
	})
}

// handleStream streams logs via SSE
func (s *Server) handleStream(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Subscribe to logs
	ch := debuglog.Subscribe()
	defer debuglog.Unsubscribe(ch)

	// Send connection established comment
	fmt.Fprintf(w, ": connected\n\n")
	flusher.Flush()

	// Handle client disconnect
	notify := r.Context().Done()

	for {
		select {
		case <-notify:
			return
		case entry := <-ch:
			data, err := json.Marshal(entry)
			if err != nil {
				continue
			}
			fmt.Fprintf(w, "data: %s\n\n", data)
			flusher.Flush()
		}
	}
}

// handleIndex serves the HTML UI
func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(htmlTemplate))
}

const htmlTemplate = `<!DOCTYPE html>
<html>
<head>
    <title>daabgo Debug Logs</title>
    <style>
        body { 
            background: #1e1e1e; 
            color: #d4d4d4; 
            font-family: 'Consolas', 'Monaco', monospace;
            padding: 20px;
            margin: 0;
            overflow: hidden;
        }
        h1 { color: #569cd6; margin-bottom: 20px; font-size: 24px; }
        .controls {
            margin-bottom: 10px;
            display: flex;
            gap: 10px;
        }
        button, input, select {
            background: #3c3c3c;
            color: #d4d4d4;
            border: 1px solid #555;
            padding: 5px 10px;
            border-radius: 3px;
        }
        button:hover { background: #4c4c4c; }
        #logs {
            background: #252526;
            border: 1px solid #3c3c3c;
            border-radius: 4px;
            padding: 15px;
            height: calc(100vh - 140px);
            overflow-y: auto;
            white-space: pre-wrap;
            word-wrap: break-word;
            font-size: 13px;
            line-height: 1.5;
        }
        .log-line { 
            margin: 2px 0; 
            display: flex;
            gap: 10px;
        }
        .log-time { color: #858585; min-width: 150px; }
        .log-level { min-width: 60px; font-weight: bold; }
        .log-debug { color: #6a9955; }
        .log-verbose { color: #6a9955; }
        .log-error { color: #f14c4c; }
        .log-info { color: #4ec9b0; }
        .status { 
            position: fixed; 
            top: 20px; 
            right: 20px; 
            padding: 5px 10px;
            border-radius: 4px;
            font-size: 12px;
        }
        .connected { background: #2d5d3a; color: #4ec9b0; }
        .disconnected { background: #5d2d2d; color: #f14c4c; }
    </style>
</head>
<body>
    <h1>🔍 daabgo Log Server</h1>
    <div id="status" class="status disconnected">Disconnected</div>
    
    <div class="controls">
        <button onclick="clearLogs()">Clear</button>
        <button onclick="toggleScroll()" id="scrollBtn">Auto-scroll: ON</button>
        <input type="text" id="filter" placeholder="Filter..." onkeyup="filterLogs()">
    </div>

    <div id="logs"></div>

    <script>
        const logsContainer = document.getElementById('logs');
        const statusEl = document.getElementById('status');
        const filterEl = document.getElementById('filter');
        let autoScroll = true;
        let eventSource;

        function connect() {
            if (eventSource) eventSource.close();

            eventSource = new EventSource('/stream');
            
            eventSource.onopen = () => {
                statusEl.textContent = 'Connected (SSE)';
                statusEl.className = 'status connected';
            };
            
            eventSource.onerror = () => {
                statusEl.textContent = 'Disconnected';
                statusEl.className = 'status disconnected';
                eventSource.close();
                setTimeout(connect, 3000);
            };
            
            eventSource.onmessage = (e) => {
                try {
                    const entry = JSON.parse(e.data);
                    addLog(entry);
                } catch (err) {
                    console.error('Parse error:', err);
                }
            };
        }
        
        function addLog(entry) {
            const line = document.createElement('div');
            line.className = 'log-line';
            
            // Level class
            let levelClass = 'log-info';
            if (entry.level === 'debug' || entry.level === 'verbose') levelClass = 'log-debug';
            if (entry.level === 'error') levelClass = 'log-error';
            
            // Format time
            const time = new Date(entry.time).toLocaleTimeString();
            
            line.innerHTML = 
                '<span class="log-time">' + time + '</span>' +
                '<span class="log-level ' + levelClass + '">' + entry.level.toUpperCase() + '</span>' +
                '<span class="log-msg ' + levelClass + '">' + escapeHtml(entry.message) + '</span>';
            
            // Filter check
            const filterText = filterEl.value.toLowerCase();
            if (filterText && !entry.message.toLowerCase().includes(filterText)) {
                line.style.display = 'none';
            }
            
            logsContainer.appendChild(line);
            
            // Limit DOM nodes
            if (logsContainer.children.length > 2000) {
                logsContainer.removeChild(logsContainer.firstChild);
            }
            
            if (autoScroll) {
                logsContainer.scrollTop = logsContainer.scrollHeight;
            }
        }
        
        function escapeHtml(text) {
            const div = document.createElement('div');
            div.textContent = text;
            return div.innerHTML;
        }

        function clearLogs() {
            logsContainer.innerHTML = '';
        }

        function toggleScroll() {
            autoScroll = !autoScroll;
            document.getElementById('scrollBtn').textContent = 'Auto-scroll: ' + (autoScroll ? 'ON' : 'OFF');
        }

        function filterLogs() {
            const filterText = filterEl.value.toLowerCase();
            const lines = logsContainer.getElementsByClassName('log-line');
            for (let line of lines) {
                const msg = line.querySelector('.log-msg').textContent;
                line.style.display = msg.toLowerCase().includes(filterText) ? '' : 'none';
            }
        }
        
        // Initial load of past logs
        fetch('/logs?limit=100').then(res => res.json()).then(data => {
            data.entries.forEach(addLog);
            connect();
        });
    </script>
</body>
</html>
`
