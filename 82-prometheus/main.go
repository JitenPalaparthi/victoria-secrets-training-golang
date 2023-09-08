package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func main() {
	recordMetrics()

	pingClicked := promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_ping_handler_click_count",
		Help: "The total number of times click pressed",
	})

	demoGuage := promauto.NewGauge(prometheus.GaugeOpts{

		Name: "myapp_ping_handler_click_guage",
		Help: "The total number of gauage click pressed",
	})

	healthClicked := promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_health_handler_click_count",
		Help: "The total number of times haealth endpoint clicked",
	})

	http.Handle("/metrics", promhttp.Handler()) // This is the key ... Prometheus is configured to fetch
	//details from this endpoint
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		pingClicked.Inc()
		demoGuage.Add(rand.Float64())
		fmt.Fprintln(w, "pong")

	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		healthClicked.Inc()
		fmt.Fprintln(w, "ok")

	})
	fmt.Println("Server started on port 2112")
	http.ListenAndServe(":2112", nil)
}
