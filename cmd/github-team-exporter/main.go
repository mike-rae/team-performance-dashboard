package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

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

	states := []string{
		"OPEN",
		"CLOSED",
		"MERGED",
	}

	for _, state := range states {

		count, err := github.PullRequestCount(
			ghClient,
			cfg.GitHubOwner,
			cfg.GitHubRepo,
			state,
		)

		if err != nil {
			log.Fatal(err)
		}

		metrics.PullRequests.WithLabelValues(
			cfg.GitHubOwner,
			cfg.GitHubRepo,
			strings.ToLower(state),
		).Set(float64(count))
	}

	// get Open PR details
	{
		openPRDetails, err := github.OpenPullRequestDetails(
			ghClient,
			cfg.GitHubOwner,
			cfg.GitHubRepo,
		)
		if err != nil {
			log.Fatal(err)
		}

		staleThresholdDays := 7
		staleCount := 0
		waitingReviewCount := 0

		for _, pr := range openPRDetails {
			ageDays := time.Since(pr.CreatedAt).Hours() / 24

			metrics.PullRequestAgeDays.WithLabelValues(
				cfg.GitHubOwner,
				cfg.GitHubRepo,
				fmt.Sprintf("%d", pr.Number),
				pr.Title,
			).Set(ageDays)

			if ageDays >= float64(staleThresholdDays) {
				staleCount++
			}

			if pr.ReviewRequestsCount > 0 && pr.ReviewsCount == 0 {
				waitingReviewCount++
			}
		}

		metrics.PullRequestsStale.WithLabelValues(
			cfg.GitHubOwner,
			cfg.GitHubRepo,
		).Set(float64(staleCount))

		metrics.PullRequestsWaitingReview.WithLabelValues(
			cfg.GitHubOwner,
			cfg.GitHubRepo,
		).Set(float64(waitingReviewCount))

		log.Printf("Open PR details collected: %d", len(openPRDetails))
		log.Printf("Stale PRs: %d", staleCount)
		log.Printf("Waiting review PRs: %d", waitingReviewCount)
	}

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
