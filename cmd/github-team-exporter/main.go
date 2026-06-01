package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.HandleFunc("/health", healthHandler)
	http.Handle("/metrics", promhttp.Handler())

	port := "2112"

	log.Printf("github-team-exporter running on http://localhost:%s", port)
	log.Printf("health endpoint:  http://localhost:%s/health", port)
	log.Printf("metrics endpoint: http://localhost:%s/metrics", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}
