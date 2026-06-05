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
	[]string{"owner", "repo", "number"},
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

var PullRequestTimeToFirstReviewHours = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "github_pull_request_time_to_first_review_hours",
		Help: "Time from pull request creation to first review in hours.",
	},
	[]string{"owner", "repo", "number"},
)

var PullRequestTimeToMergeHours = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "github_pull_request_time_to_merge_hours",
		Help: "Time from pull request creation to merge in hours.",
	},
	[]string{"owner", "repo", "number"},
)

var PullRequestsWithoutReviewers = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "github_pull_requests_without_reviewers",
		Help: "Number of open pull requests without reviewers assigned.",
	},
	[]string{"owner", "repo"},
)

var PullRequestsReviewBacklog = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "github_pull_requests_review_backlog",
		Help: "Number of open pull requests with no submitted reviews.",
	},
	[]string{"owner", "repo"},
)

func Register() {
	prometheus.MustRegister(PullRequests)
	prometheus.MustRegister(PullRequestAgeDays)
	prometheus.MustRegister(PullRequestsStale)
	prometheus.MustRegister(PullRequestsWaitingReview)
	prometheus.MustRegister(PullRequestTimeToFirstReviewHours)
	prometheus.MustRegister(PullRequestTimeToMergeHours)
	prometheus.MustRegister(PullRequestsWithoutReviewers)
	prometheus.MustRegister(PullRequestsReviewBacklog)
}
