package main

import (
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var DB *gorm.DB

func main() {

	DB = connectWithGorm()
	http.HandleFunc("/point", point)
	http.HandleFunc("/status", status)
	http.HandleFunc("/start", start)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func point(w http.ResponseWriter, r *http.Request) {
	result := map[string]interface{}{}
	DB.Table("exchange_models").Where("exchange_id = ?", "btc").Take(&result)
	d, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ACTIVE CONNECTIONS: ", d.Stats().OpenConnections)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println(err)
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	database, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(
		"Open: ", database.Stats().OpenConnections,
		" Idle: ", database.Stats().Idle,
		" Dead: ", database.Stats().MaxIdleTimeClosed,
	)
}

func start(w http.ResponseWriter, r *http.Request) {
	i := 0
	for i < 100 {
		go func() {
			response, err := http.NewRequest(http.MethodPost, "http://localhost:8080/point", nil)
			if err != nil {
				log.Println("err: ", err)
			}
			client := &http.Client{}
			resp, err := client.Do(response)
			if err != nil {
				log.Println(err)
			}
			defer resp.Body.Close()
		}()
		i++
	}

}
