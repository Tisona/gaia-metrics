package main

import (
	"fmt"
	"log"
	"metrics/metrics"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var APP_HOST = "127.0.0.1:3000"

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
	fmt.Println("listening on", APP_HOST)

	http.HandleFunc("/", index)

	err := http.ListenAndServe(APP_HOST, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
