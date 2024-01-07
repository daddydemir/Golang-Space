package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

var shutdownCh chan struct{}
var wg sync.WaitGroup

func main() {

	shutdownCh = make(chan struct{})

	http.HandleFunc("/buyTheProduct", buyTheProduct)
	http.HandleFunc("/consume", startConsume)

	srv := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	select {
	case sig := <-signalCh:
		log.Printf("Received signal: %v. Shutting down...", sig)
		close(shutdownCh)
		wg.Wait()
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("Server shutdown returned an error: %v\n", err)
		}
	}

}

func startConsume(w http.ResponseWriter, r *http.Request) {

	log.Println("Consumer is started.")
	wg.Add(1)
	go Consume()

}

func Consume() {

	defer wg.Done()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "gracefulShutdown",
		GroupID:     "group-1",
		StartOffset: kafka.FirstOffset,
	})

	for {
		select {
		case <-shutdownCh:
			log.Println("kafka consumer closing.")
			reader.Close()
			return

		default:
			m, err := reader.FetchMessage(context.Background())
			if err != nil {
				log.Println("read error: ", err)
			}

			log.Printf("message: %v ", string(m.Value))
			time.Sleep(2 * time.Second)

			err = reader.CommitMessages(context.Background(), m)
			if err != nil {
				log.Println("commit error: ", err)
			} else {
				log.Println("commit successfully : ", string(m.Value))
			}
		}

	}

}

func buyTheProduct(w http.ResponseWriter, r *http.Request) {

	// insert orders
	saveOrder()

	// update unit stock
	updateStock()

	// notify to seller

	// notify to customer

	log.Println("the payoff successfully.")
}

func saveOrder() {
	time.Sleep(2 * time.Second)
	log.Println("Order saved.")
}

func updateStock() {
	time.Sleep(2 * time.Second)
	log.Println("Stock updated.")
}
