package controllers

import (
	"log"
	"net/http"
	ws "scm-api/ws"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebSocketHandler(broker *ws.Broker) echo.HandlerFunc {
	return func(c echo.Context) error {
		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}

		client := &ws.Client{
			RoomID: c.QueryParam("room"),
			Conn:   conn,
			Send:   make(chan []byte, 256),
		}

		broker.Register <- client

		go handleClient(client, broker)

		log.Printf("Client connected to room: %s", client.RoomID)

		return nil
	}
}

func handleClient(client *ws.Client, broker *ws.Broker) {
	defer func() {
		broker.Unregister <- client
		client.Conn.Close()
	}()

	// This goroutine reads messages from the client.Send channel
	// and writes them to the WebSocket connection.
	go func() {
		for message := range client.Send {
			if err := client.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("Error writing message: %v", err)
				break // Exit the loop if there's an error writing to the WebSocket
			}
		}
	}()

	// This loop reads messages from the WebSocket connection
	for {
		_, msg, err := client.Conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break // Exit the loop if there's an error reading from the WebSocket
		}
		broker.Broadcast <- ws.Message{RoomID: client.RoomID, Message: msg}
	}
}
