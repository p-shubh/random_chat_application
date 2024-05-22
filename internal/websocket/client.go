// /internal/websocket/client.go
package websockets

import (
    "github.com/gorilla/websocket"
    "github.com/google/uuid"
    "log"
)

type Client struct {
    ID   uuid.UUID
    Conn *websocket.Conn
    Send chan []byte
    Hub  *Hub
}

func (c *Client) ReadMessages() {
    defer func() {
        c.Hub.Unregister <- c
        c.Conn.Close()
    }()

    for {
        _, message, err := c.Conn.ReadMessage()
        if err != nil {
            log.Printf("error: %v", err)
            break
        }
        c.Hub.Broadcast <- message
    }
}

func (c *Client) WriteMessages() {
    defer c.Conn.Close()
    for {
        select {
        case message, ok := <-c.Send:
            if !ok {
                c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }

            c.Conn.WriteMessage(websocket.TextMessage, message)
        }
    }
}
