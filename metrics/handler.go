package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "endpoint", "status"},
		)
)

func handler() {
	
	http.Handle("/metrics", promhttp.Handler())
	
	http.HandleFunc("/", root)
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(" error : " , err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	requestsTotal.WithLabelValues(r.Method, r.URL.Path, "200")
	
	w.WriteHeader(http.StatusOK)
	_ , err := w.Write([]byte("Hello World!"))
	if err != nil {
		log.Println("/root error: ", err)	
	}
}