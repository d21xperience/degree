package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WebSocketServer struct {
	clients   map[*websocket.Conn]bool
	broadcast chan interface{}
	upgrader  websocket.Upgrader
	mutex     sync.RWMutex
	isRunning bool
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan interface{}, 100),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // Allow all origins for development
			},
		},
		isRunning: false,
	}
}

func (w *WebSocketServer) Start(port string) error {
	if w.isRunning {
		return status.Error(codes.AlreadyExists, "WebSocket server already running")
	}

	http.HandleFunc("/ws", w.handleWebSocket)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	go func() {
		log.Printf("WebSocket server starting on port %s", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Printf("WebSocket server error: %v", err)
		}
	}()

	go w.broadcastMessages()
	w.isRunning = true

	return nil
}

func (w *WebSocketServer) handleWebSocket(writer http.ResponseWriter, request *http.Request) {
	conn, err := w.upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	// Register client
	w.mutex.Lock()
	w.clients[conn] = true
	clientCount := len(w.clients)
	w.mutex.Unlock()

	log.Printf("New WebSocket client connected. Total clients: %d", clientCount)

	// Send connection success message
	w.sendMessage(conn, map[string]interface{}{
		"type": "connection_status",
		"data": map[string]interface{}{
			"status":  "connected",
			"message": "Successfully connected to blockchain monitor",
		},
	})

	// Handle incoming messages (optional)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Client disconnected: %v", err)
			break
		}
	}

	// Unregister client
	w.mutex.Lock()
	delete(w.clients, conn)
	clientCount = len(w.clients)
	w.mutex.Unlock()
	log.Printf("Client disconnected. Total clients: %d", clientCount)
}

func (w *WebSocketServer) broadcastMessages() {
	for {
		message := <-w.broadcast
		w.mutex.RLock()
		for client := range w.clients {
			err := w.sendMessage(client, message)
			if err != nil {
				// Remove dead connections
				delete(w.clients, client)
				client.Close()
			}
		}
		w.mutex.RUnlock()
	}
}

func (w *WebSocketServer) sendMessage(conn *websocket.Conn, message interface{}) error {
	return conn.WriteJSON(message)
}

func (w *WebSocketServer) BroadcastMessage(message interface{}) {
	select {
	case w.broadcast <- message:
	default:
		log.Println("WebSocket broadcast channel is full, dropping message")
	}
}

func (w *WebSocketServer) GetClientCount() int {
	w.mutex.RLock()
	defer w.mutex.RUnlock()
	return len(w.clients)
}

func (w *WebSocketServer) IsRunning() bool {
	return w.isRunning
}
