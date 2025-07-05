package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/cancelledbit/smartctl_ssacli_exporter/exporter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	numerationType = flag.String("numerationType", "default", "numeration type, default or from_one")
	listenAddr     = flag.String("listen", ":9633", "address for exporter")
	metricsPath    = flag.String("path", "/metrics", "URL path for surfacing collected metrics")
)

func main() {
	flag.Parse()
	if numerationType == nil {
		panic("numeration type is required")
	}
	prometheus.MustRegister(exporter.New(exporter.NumerationType(*numerationType)))

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, *metricsPath, http.StatusMovedPermanently)
	})

	log.Println("Beginning to serve on port ", *listenAddr)

	if err := http.ListenAndServe(*listenAddr, nil); err != nil {
		log.Fatalf("Cannot start exporter: %s", err)
	}
}
