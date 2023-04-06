package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

var PORT = ":1234"
var counter = prometheus.NewCounter(prometheus.CounterOpts{
	Namespace: "mattchang",
	Name:      "My_Counter",
	Help:      "This is my counter",
})

var gauge = prometheus.NewGauge(prometheus.GaugeOpts{
	Namespace: "mattchang",
	Name:      "My_Gauge",
	Help:      "This is my gauge",
})

var histogram = prometheus.NewHistogram(prometheus.HistogramOpts{
	Namespace: "mattchang",
	Name:      "My_Histogram",
	Help:      "This is my histogram",
})

var summary = prometheus.NewSummary(prometheus.SummaryOpts{
	Namespace: "mattchang",
	Name:      "My_Summary",
	Help:      "This is my summary",
})

func main() {
	rand.Seed(time.Now().Unix())

	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)

	go func() {
		for {
			counter.Add(rand.Float64() * 5)
			gauge.Add(rand.Float64()*15 - 5)
			histogram.Observe(rand.Float64() * 10)
			summary.Observe(rand.Float64() * 10)
			time.Sleep(2 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Listening to port", PORT)
	fmt.Println(http.ListenAndServe(PORT, nil))
}
