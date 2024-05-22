// /internal/handlers/chat_handler.go
package handlers

import (
	"net/http"
	websockets "rca/internal/websocket"

	// "rca/internal/websocket"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.Error(c.Writer, "Could not open websocket connection", http.StatusBadRequest)
		return
	}

	client := &websockets.Client{ID: uuid.New(), Conn: conn, Send: make(chan []byte)}

	client.Hub.Register <- client

	go client.ReadMessages()
	go client.WriteMessages()
}
