package ws

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// ===================================================
// Declaring types here to avoid external usage issues
// ===================================================

// Client represents a single WebSocket connection. Each client is
// identified by an ID and can be subscribed to a room to receive messages.
type Client struct {
	ID     string          // Unique identifier for the client
	RoomID string          // ID of the room the client is subscribed to
	Conn   *websocket.Conn // WebSocket connection for this client
	Send   chan []byte     // Channel for sending messages to this client
}

// Broker manages all active clients and rooms. It facilitates registering
// and unregistering clients, and broadcasting messages to rooms.
type Broker struct {
	Clients    map[*Client]bool            // Tracks all connected clients
	Rooms      map[string]map[*Client]bool // Maps room IDs to clients subscribed to that room
	Register   chan *Client                // Channel for requests to register new clients
	Unregister chan *Client                // Channel for requests to unregister clients
	Broadcast  chan Message                // Channel for broadcasting messages to rooms
	Mu         sync.Mutex                  // Mutex to ensure thread-safe access to Clients and Rooms
}

// Message represents a message to be sent to a room. It contains the
// room ID and the message content.
type Message struct {
	RoomID  string `json:"room_id"`
	Message []byte
}

// ===================================================
// ***************************************************
// ===================================================

func NewBroker() *Broker {
	return &Broker{
		Clients:    make(map[*Client]bool),
		Rooms:      make(map[string]map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Message),
	}
}

// Run starts the broker's main loop, listening for register, unregister,
// and broadcast requests, and handling them appropriately.
func (b *Broker) Run() {
	for {
		select {
		case client := <-b.Register:
			// Lock the broker to safely update the Clients and Rooms maps
			b.Mu.Lock()
			// Register the new client
			b.Clients[client] = true
			// Add the client to the specified room. If the room doesn't exist, create it.
			if _, ok := b.Rooms[client.RoomID]; !ok {
				b.Rooms[client.RoomID] = make(map[*Client]bool)
			}
			b.Rooms[client.RoomID][client] = true
			// Unlock the broker after updating
			b.Mu.Unlock()
		case client := <-b.Unregister:
			// Lock the broker to safely update the Clients and Rooms maps
			b.Mu.Lock()
			// Unregister the client if it exists
			if _, ok := b.Clients[client]; ok {
				// Remove the client from the global list and its room
				delete(b.Clients, client)
				delete(b.Rooms[client.RoomID], client)
				// Close the client's send channel to stop sending messages
				close(client.Send)
			}
			// Unlock the broker after updating
			b.Mu.Unlock()
		case message := <-b.Broadcast:
			b.Mu.Lock()
			log.Printf("Broadcasting message to room: %s", message.RoomID)
			if clients, ok := b.Rooms[message.RoomID]; ok {
				for client := range clients {
					select {
					case client.Send <- message.Message:
					default:
						close(client.Send)
						delete(b.Clients, client)
						delete(b.Rooms[message.RoomID], client)
					}
				}
			}
			b.Mu.Unlock()
		}
	}
}
