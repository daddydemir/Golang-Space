package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var stock = 2

var ws *websocket.Conn

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("conn.ReadMessage : ", err)
		}
		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("conn.WriteMessage : ", err)
			return
		}
	}
}

func socket(w http.ResponseWriter, r *http.Request) {
	ws, _ = upgrader.Upgrade(w, r, nil)
	/*
		err := ws.WriteMessage(1, []byte("Hello World!"))
		if err != nil {
			log.Println("ws.WriteMessage : ", err)
		}
		reader(ws)

	*/
}

func reduceStock(w http.ResponseWriter, r *http.Request) {

	if stock <= 0 {
		_ = ws.WriteMessage(websocket.TextMessage, []byte("the stock is empty."))
		_ = json.NewEncoder(w).Encode("the stock already empty.")
	} else {
		stock = stock - 1
		_ = json.NewEncoder(w).Encode("the stock has been reduced by one.")
	}
}

func main() {

	http.HandleFunc("/socket", socket)
	http.HandleFunc("/reduceStock", reduceStock)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
