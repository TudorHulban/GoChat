package main

import (
	"github.com/gofiber/websocket/v2"
)

// User is used for variables injecting data about user.
type User struct {
	Name     string
	socket   websocket.Conn
	commSend chan []byte // used to send messages
	room     *Room       // room where user is chatting
}

func (u *User) Read() {
	defer u.socket.Close()

	for {
		_, msg, errRead := u.socket.ReadMessage()
		if errRead != nil {
			return
		}

		u.room.commBroadcast <- msg
	}
}

func (u *User) Write() {
	defer u.socket.Close()

	for msg := range u.commSend {
		if errSend := u.socket.WriteMessage(websocket.TextMessage, msg); errSend != nil {
			return
		}
	}
}
