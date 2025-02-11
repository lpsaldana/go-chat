package models

import "github.com/gorilla/websocket"

type User struct {
	Conn *websocket.Conn
	Send chan []byte
}

func (c *User) WritePump() {
	defer c.Conn.Close()
	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}
