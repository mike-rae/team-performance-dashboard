package metrics

import "github.com/prometheus/client_golang/prometheus"

var PullRequests = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "github_pull_requests",
		Help: "Number of GitHub pull requests by state.",
	},
	[]string{"owner", "repo", "state"},
)

func Register() {
	prometheus.MustRegister(PullRequests)
}
