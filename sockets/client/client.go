package main

import (
	"log"
	"net"
)

func main() {

	connectToServer()
}

func connectToServer() {

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	_, err = conn.Write([]byte("Hello, world!"))
	if err != nil {
		log.Println(err)
		return
	}

	buffer := make([]byte, 1024)
	for {
		_, err = conn.Read(buffer)
		if err != nil {
			log.Println(err)
		}
		log.Println("Client: ", string(buffer))
		break
	}

	conn.Close()
}
