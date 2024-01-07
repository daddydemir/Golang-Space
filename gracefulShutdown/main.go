package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	http.HandleFunc("/buyTheProduct", buyTheProduct)

	srv := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	log.Println("got interruption signal.")
	if err := srv.Shutdown(context.TODO()); err != nil {
		log.Printf("server shutdown returned an err: %v\n", err)
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
