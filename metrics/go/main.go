package main

import (
	"app/metrics"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

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
	fmt.Println("listening ...")

	http.HandleFunc("/", index)

	err := http.ListenAndServe("127.0.0.1:3000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
