package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/stat", stat)
	http.HandleFunc("/ping", ping)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Panic(err)
	}
}
