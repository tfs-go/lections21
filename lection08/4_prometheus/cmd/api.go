package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	reg := prometheus.NewRegistry()
	wrappedReg := prometheus.WrapRegistererWith(prometheus.Labels{"project_name": "test_prom"}, reg)
	wrappedReg.MustRegister(
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		collectors.NewGoCollector(),
	)

	metricTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "some_total",
		Help: "Some total",
	}, []string{"tag"})

	reg.MustRegister(metricTotal)

	go func() {
		for {
			metricTotal.WithLabelValues("tag1").Add(float64(rand.Intn(10)))
			metricTotal.WithLabelValues("tag2").Add(float64(rand.Intn(30)))
			time.Sleep(time.Second)
		}
	}()

	router := chi.NewRouter()
	router.Mount("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	httpServer := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	_ = httpServer.ListenAndServe()
}
