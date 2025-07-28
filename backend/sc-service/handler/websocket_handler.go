package handlers

// import (
// 	"log"
// 	"net/http"
// 	"sc-service/services/clients"
// 	"sync"
// 	"time"

// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/gorilla/websocket"
// )

// type WebSocketHandler struct {
// 	// ethereumService *services.EthereumService
// 	clients   map[*websocket.Conn]bool
// 	broadcast chan interface{}
// 	mutex     sync.RWMutex
// }

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true // Allow all origins for development
// 	},
// }

// func NewWebSocketHandler() *WebSocketHandler {
// 	handler := &WebSocketHandler{
// 		// ethereumService: ethereumService,
// 		clients:   make(map[*websocket.Conn]bool),
// 		broadcast: make(chan interface{}, 100),
// 	}

// 	go handler.broadcastMessages()
// 	go handler.monitorNetwork()

// 	return handler
// }

// func (h *WebSocketHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Printf("Failed to upgrade connection: %v", err)
// 		return
// 	}
// 	defer conn.Close()

// 	// Register client
// 	h.mutex.Lock()
// 	h.clients[conn] = true
// 	h.mutex.Unlock()

// 	log.Printf("New WebSocket client connected. Total clients: %d", len(h.clients))

// 	// Send initial network info
// 	if info, err := clients.Bloc.GetNetworkInfo(); err == nil {
// 		h.sendMessage(conn, map[string]interface{}{
// 			"type": "initial_info",
// 			"data": info,
// 		})
// 	}

// 	// Handle incoming messages (optional)
// 	for {
// 		_, _, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Printf("Client disconnected: %v", err)
// 			break
// 		}
// 	}

// 	// Unregister client
// 	h.mutex.Lock()
// 	delete(h.clients, conn)
// 	h.mutex.Unlock()
// 	log.Printf("Client disconnected. Total clients: %d", len(h.clients))
// }

// func (h *WebSocketHandler) broadcastMessages() {
// 	for {
// 		message := <-h.broadcast
// 		h.mutex.RLock()
// 		for client := range h.clients {
// 			err := h.sendMessage(client, message)
// 			if err != nil {
// 				// Remove dead connections
// 				delete(h.clients, client)
// 				client.Close()
// 			}
// 		}
// 		h.mutex.RUnlock()
// 	}
// }

// func (h *WebSocketHandler) sendMessage(conn *websocket.Conn, message interface{}) error {
// 	return conn.WriteJSON(message)
// }

// func (h *WebSocketHandler) monitorNetwork() {
// 	// Subscribe to new blocks
// 	headers := make(chan *types.Header)
// 	unsubscribe, err := h.ethereumService.SubscribeNewHeads(headers)
// 	if err != nil {
// 		log.Printf("Failed to subscribe to new heads: %v", err)
// 		return
// 	}
// 	defer unsubscribe()

// 	// Send periodic network info updates
// 	ticker := time.NewTicker(10 * time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case header := <-headers:
// 			// Send block update
// 			h.broadcast <- map[string]interface{}{
// 				"type": "new_block",
// 				"data": map[string]interface{}{
// 					"block_number": header.Number.Uint64(),
// 					"block_hash":   header.Hash().Hex(),
// 					"timestamp":    header.Time,
// 				},
// 			}

// 		case <-ticker.C:
// 			// Send periodic network info
// 			if info, err := h.ethereumService.GetNetworkInfo(); err == nil {
// 				h.broadcast <- map[string]interface{}{
// 					"type": "network_info",
// 					"data": info,
// 				}
// 			} else {
// 				log.Printf("Failed to get network info: %v", err)
// 			}
// 		}
// 	}
// }
