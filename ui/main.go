package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var messageFromCLI string
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)
var upgrader = websocket.Upgrader{}
var html = template.Must(template.ParseFiles("./ui/index.html"))

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		err := html.Execute(c.Writer, nil)
		if err != nil {
			fmt.Println(err)
		}
	})

	r.GET("/ws", func(c *gin.Context) {
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
		}
		defer ws.Close()

		clients[ws] = true

		for {
			_, _, err := ws.ReadMessage()
			if err != nil {
				delete(clients, ws)
				break
			}
		}
	})

	r.POST("/data", func(c *gin.Context) {
		var json struct {
			Message string `json:"message"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		messageFromCLI = json.Message
		fmt.Println("Received message from CLI:", json.Message)
		broadcast <- json.Message
		c.JSON(http.StatusOK, gin.H{"status": "Message received"})
	})

	go handleMessages()

	r.Run(":8080")
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			message := struct {
				Message string `json:"message"`
			}{
				Message: msg,
			}
			data, _ := json.Marshal(message)
			err := client.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				_ = client.Close()
				delete(clients, client)
			}
		}
	}
}
