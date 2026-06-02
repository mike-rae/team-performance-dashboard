package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mike-rae/engineering-observability-dashboard/internal/config"
	"github.com/mike-rae/engineering-observability-dashboard/internal/github"
	"github.com/mike-rae/engineering-observability-dashboard/internal/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	cfg := config.Load()

	metrics.Register()

	log.Printf(
		"Loaded config for %s/%s",
		cfg.GitHubOwner,
		cfg.GitHubRepo,
	)

	ghClient := github.NewClient(cfg.GitHubToken)

	openPRs, err := github.OpenPullRequestCount(
		ghClient,
		cfg.GitHubOwner,
		cfg.GitHubRepo,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Open PRs for %s/%s: %d", cfg.GitHubOwner, cfg.GitHubRepo, openPRs)

	metrics.OpenPullRequests.WithLabelValues(
		cfg.GitHubOwner,
		cfg.GitHubRepo,
	).Set(float64(openPRs))

	http.HandleFunc("/health", healthHandler)
	http.Handle("/metrics", promhttp.Handler())

	port := "2112"

	log.Printf("github-team-exporter running on http://localhost:%s", port)
	log.Printf("health endpoint:  http://localhost:%s/health", port)
	log.Printf("metrics endpoint: http://localhost:%s/metrics", port)

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthHandler(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintln(writer, "OK")
}
