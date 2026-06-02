package metrics

import "github.com/prometheus/client_golang/prometheus"

var OpenPullRequests = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "github_open_pull_requests",
		Help: "Number of currently open GitHub pull requests.",
	},
	[]string{"owner", "repo"},
)

func Register() {
	prometheus.MustRegister(OpenPullRequests)
}
