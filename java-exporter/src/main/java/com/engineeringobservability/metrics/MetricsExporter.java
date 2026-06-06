package com.engineeringobservability.metrics;

import io.prometheus.metrics.core.metrics.Gauge;

public class MetricsExporter {

    public static final Gauge pullRequests = Gauge.builder()
            .name("java_github_pull_requests")
            .help("Number of GitHub pull requests by state.")
            .labelNames("owner", "repo", "state")
            .register();

    public static final Gauge pullRequestAgeDays = Gauge.builder()
            .name("java_github_pull_request_age_days")
            .help("Age of open GitHub pull requests in days.")
            .labelNames("owner", "repo", "number")
            .register();

    public static final Gauge pullRequestsStale = Gauge.builder()
            .name("java_github_pull_requests_stale")
            .help("Number of open GitHub pull requests older than the stale threshold.")
            .labelNames("owner", "repo")
            .register();
}
