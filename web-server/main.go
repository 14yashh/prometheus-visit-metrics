package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests for this application.",
	})
)

func homePage(w http.ResponseWriter, r *http.Request) {
	httpRequestsTotal.Inc()

	w.Write([]byte("hey! our server is running."))
}

func main() {
	http.HandleFunc("/", homePage)

	http.Handle("/metrics", promhttp.Handler())

	log.Println("Go server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
