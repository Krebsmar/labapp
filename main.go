package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const Version = "1.0.5"

var (
	versionGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "app_version",
			Help: "Current version of the application.",
		},
		[]string{"version"},
	)
)

func main() {
	log.Println("Starting the application...")
	log.Printf("Version: %s\n", Version)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "LabApp Version: %s\n", Version)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	// Start the load generator
	go generateLoad()

	http.Handle("/metrics", promhttp.Handler())

	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
	prometheus.MustRegister(versionGauge)

	versionGauge.WithLabelValues(Version).Set(1)

	http.ListenAndServe(":8080", nil)
}

func generateLoad() {
	for {
		// Perform some computation
		for i := 0; i < 1e6; i++ {
			_ = i * i
			if i%100000 == 0 { // Log every 100000 iterations
				log.Printf("Computation iteration: %d\n", i)
			}
		}

		// Sleep for a while to avoid consuming too much CPU
		time.Sleep(100 * time.Millisecond)
	}
}
