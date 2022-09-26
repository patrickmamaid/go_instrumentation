package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "pmapp_infinite_ops_2s_total",
		Help: "The total number of processed events",
	})
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	recordMetrics() // infinite counter

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
	fmt.Printf("Server started")
}


