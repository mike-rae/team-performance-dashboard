package metrics

import "github.com/prometheus/client_golang/prometheus"

var PullRequests = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "github_pull_requests",
		Help: "Number of GitHub pull requests by state.",
	},
	[]string{"owner", "repo", "state"},
)

var PullRequestAgeDays = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "github_pull_request_age_days",
		Help: "Age of open GitHub pull requests in days.",
	},
	[]string{"owner", "repo", "number", "title"},
)

var PullRequestsStale = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "github_pull_requests_stale",
		Help: "Number of open GitHub pull requests older than the stale threshold.",
	},
	[]string{"owner", "repo"},
)

var PullRequestsWaitingReview = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "github_pull_requests_waiting_review",
		Help: "Number of open GitHub pull requests waiting for review.",
	},
	[]string{"owner", "repo"},
)

func Register() {
	prometheus.MustRegister(PullRequests)
	prometheus.MustRegister(PullRequestAgeDays)
	prometheus.MustRegister(PullRequestsStale)
	prometheus.MustRegister(PullRequestsWaitingReview)
}
