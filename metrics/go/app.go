package main

import (
	"fmt"
	"log"
	"metrics/metrics"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// var APP_HOST = "127.0.0.1:9000"

func index(w http.ResponseWriter, r *http.Request) {
	latest_block_height, desync, n_peers := metrics.CollectMetrics()

	registry := metrics.RegisterMetrics(latest_block_height, desync, n_peers)

	handler := promhttp.HandlerFor(
		registry,
		promhttp.HandlerOpts{},
	)

	handler.ServeHTTP(w, r)
}

func main() {
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")

	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "9000"
	}

	address := fmt.Sprintf("%s:%s", host, port)

	fmt.Println("listening on", address)

	http.HandleFunc("/metrics", index)

	err := http.ListenAndServe(address, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
