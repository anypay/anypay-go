package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/anypay/anypay-go/log"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

// Client represents a WebSocket client
type Client struct {
	conn *websocket.Conn
	send chan []byte
}

// Hub maintains active clients and broadcasts messages to clients
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
	lock       sync.Mutex
}

// Message represents a message structure
type Message struct {
	Topic   string          `json:"topic"`
	Payload json.RawMessage `json:"payload"`
}

var hub = Hub{
	broadcast:  make(chan Message),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message.Payload:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (c *Client) readPump() {
	defer func() {
		hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				// Log error
			}
			break
		}
		// log the message
		fmt.Println(string(message))
		log.Log.Info("websocket.message", zap.String("message", string(message)))

	}
}

func (c *Client) writePump() {
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// The hub closed the channel
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func serveWs(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// Log error
		return
	}
	client := &Client{conn: conn, send: make(chan []byte, 256)}
	hub.register <- client

	initialMessage := Message{
		Topic:   "websocket.connected",
		Payload: json.RawMessage(`{"success": true}`),
	}
	messageBytes, err := json.Marshal(initialMessage)
	if err != nil {
		// Handle error
		return
	}
	client.send <- messageBytes

	go client.writePump()
	go client.readPump()
}

func SetupWebsocketRoutes(router *gin.Engine) {
	router.GET("/ws", func(c *gin.Context) {
		serveWs(c)
	})
}

func init() {
	go hub.run()
}
