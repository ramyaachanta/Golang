package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Order)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()
	clients[conn] = true

	for {
		var order Order
		err := conn.ReadJSON(&order)
		if err != nil {
			delete(clients, conn)
			break
		}
		broadcast <- order
	}
}

func broadcastOrders() {
	for {
		order := <-broadcast
		for client := range clients {
			err := client.WriteJSON(order)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}
