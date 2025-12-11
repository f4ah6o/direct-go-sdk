// Debug server that streams logs via WebSocket
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var (
	clients   = make(map[*websocket.Conn]bool)
	clientsMu sync.Mutex
	broadcast = make(chan string, 100)
)

func main() {
	// WebSocket endpoint for log streaming
	http.HandleFunc("/ws", handleWS)

	// HTTP endpoint to receive logs
	http.HandleFunc("/log", handleLog)

	// Simple HTML page to view logs
	http.HandleFunc("/", handleIndex)

	go broadcaster()

	port := ":9999"
	fmt.Printf("Debug server running at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	clientsMu.Lock()
	clients[conn] = true
	clientsMu.Unlock()

	defer func() {
		clientsMu.Lock()
		delete(clients, conn)
		clientsMu.Unlock()
	}()

	// Keep connection alive
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

func handleLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}

	// Read log message
	buf := make([]byte, 10240)
	n, _ := r.Body.Read(buf)
	msg := string(buf[:n])

	// Broadcast to all WebSocket clients
	broadcast <- msg

	w.WriteHeader(http.StatusOK)
}

func broadcaster() {
	for msg := range broadcast {
		clientsMu.Lock()
		for conn := range clients {
			err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				conn.Close()
				delete(clients, conn)
			}
		}
		clientsMu.Unlock()
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
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
        }
        h1 { color: #569cd6; margin-bottom: 20px; }
        #logs {
            background: #252526;
            border: 1px solid #3c3c3c;
            border-radius: 4px;
            padding: 15px;
            height: calc(100vh - 120px);
            overflow-y: auto;
            white-space: pre-wrap;
            word-wrap: break-word;
            font-size: 13px;
            line-height: 1.5;
        }
        .log-line { margin: 2px 0; }
        .log-debug { color: #6a9955; }
        .log-error { color: #f14c4c; }
        .log-info { color: #4ec9b0; }
        .status { 
            position: fixed; 
            top: 10px; 
            right: 20px; 
            padding: 5px 10px;
            border-radius: 4px;
        }
        .connected { background: #2d5d3a; color: #4ec9b0; }
        .disconnected { background: #5d2d2d; color: #f14c4c; }
    </style>
</head>
<body>
    <h1>🔍 daabgo Debug Logs</h1>
    <div id="status" class="status disconnected">Disconnected</div>
    <div id="logs"></div>
    <script>
        const logs = document.getElementById('logs');
        const status = document.getElementById('status');
        let ws;
        
        function connect() {
            ws = new WebSocket('ws://' + location.host + '/ws');
            
            ws.onopen = () => {
                status.textContent = 'Connected';
                status.className = 'status connected';
            };
            
            ws.onclose = () => {
                status.textContent = 'Disconnected';
                status.className = 'status disconnected';
                setTimeout(connect, 2000);
            };
            
            ws.onmessage = (e) => {
                const line = document.createElement('div');
                line.className = 'log-line';
                const text = e.data;
                
                if (text.includes('[DEBUG]')) {
                    line.classList.add('log-debug');
                } else if (text.includes('error') || text.includes('Error')) {
                    line.classList.add('log-error');
                } else {
                    line.classList.add('log-info');
                }
                
                line.textContent = text;
                logs.appendChild(line);
                logs.scrollTop = logs.scrollHeight;
            };
        }
        
        connect();
    </script>
</body>
</html>`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
