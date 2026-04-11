package main

import (
	"net/http"
	"time"

	"github.com/gookit/goutil/dump"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main01() {

	temp := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "test_temp_gauge",
		Help: "this is a test gauge",
	})
	prometheus.MustRegister(temp)

	var i int
	go func() {
		for {
			i++
			if i%2 == 0 {
				temp.Inc()
			}
			time.Sleep(time.Second)
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8099", nil)
	if err != nil {
		dump.P("ListenAndServe err", err)
	}
}
