package main

import (
	"log"
	"net"
)

func main() {

	StartServer()
}

func StartTcpServer() {

	connection, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	_, err = connection.Write([]byte("hello world"))
	if err != nil {
		log.Println(err)
	}

	buffer := make([]byte, 1024)
	length, err := connection.Read(buffer)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(buffer[:length]))

	defer connection.Close()
}

func StartServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = conn.Write([]byte("hello client..."))
	if err != nil {
		log.Println(err)
	}
	
	log.Println("Data : ", string(buf))
}
