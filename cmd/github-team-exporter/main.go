package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mike-rae/engineering-observability-dashboard/internal/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	cfg := config.Load()

	log.Printf(
		"Loaded config for %s/%s",
		cfg.GitHubOwner,
		cfg.GitHubRepo,
	)

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

func healthHandler(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintln(writer, "OK")
}
