# Engineering Observability Dashboard

A simple engineering observability platform built using Go, GitHub GraphQL, Prometheus and Grafana.

## Dashboard Preview

![Engineering Observability Dashboard](docs/screenshots/milestone-4/grafana-dashboard-overview.png)

## Overview

Engineering teams generate large amounts of delivery data through source control systems, pull requests, code reviews, issue tracking platforms, and CI/CD pipelines.

The goal of this project is to explore whether useful engineering observability insights can be derived from GitHub using open-source tooling and a lightweight architecture.

This project focuses on understanding engineering flow and delivery health rather than measuring individual productivity.

## Project Goals

The long-term goal is to provide visibility into metrics such as:

* Open pull requests
* Stale pull requests
* Pull request age
* Review activity
* Pull request throughput
* Time to merge
* Delivery bottlenecks
* Engineering flow trends

## Target Architecture

```text
GitHub GraphQL API
        ↓
Go Exporter
        ↓
Prometheus
        ↓
Grafana
```

## Current Status

### Milestone 1 – Foundations Complete ✅

Completed:

- Repository structure established
- Go module initialised
- Local development environment configured
- Health endpoint implemented
- Metrics endpoint implemented

### Milestone 2 – GitHub GraphQL Integration Complete ✅

Completed:

- Environment configuration loading
- GitHub GraphQL client
- GitHub authentication using personal access token
- Pull request count query
- Custom Prometheus metrics
- Pull request metrics by state

### Milestone 3 – Docker & Prometheus Integration Complete ✅

Completed:

- Dockerfile implementation
- Docker Compose environment
- Prometheus configuration
- Exporter containerisation
- Prometheus metric scraping
- End-to-end observability pipeline

### Milestone 4 – Grafana Dashboard Complete ✅

Completed:

- Grafana integration
- Prometheus datasource configuration
- Engineering observability dashboard
- Dashboard thresholds and KPI visualisation
- Pull request flow metrics visualisation

## Current Architecture

```text
GitHub GraphQL API
        ↓
Go Exporter
        ↓
Prometheus
        ↓
Grafana
```

## Next Target Architecture

```text
GitHub GraphQL API
        ↓
Go Exporter
        ↓
Prometheus
        ↓
Grafana
        ↓
Engineering Flow Analytics
```

## Current Functionality

The exporter currently provides:

* `/health`
* `/metrics`
* GitHub pull request metrics by state

Current custom metric:
```text
github_pull_requests{owner="mike-rae",repo="engineering-observability-dashboard",state="open"}

github_pull_requests{owner="mike-rae",repo="engineering-observability-dashboard",state="closed"}

github_pull_requests{owner="mike-rae",repo="engineering-observability-dashboard",state="merged"}
```

The project does not yet provide:
- Multi-repository support
- Historical trend analysis
- Time-to-review metrics
- Time-to-merge metrics
- Review turnaround metrics

---

## Getting Started

### Clone the Repository

```bash
git clone git@github.com:mike-rae/engineering-observability-dashboard.git
cd engineering-observability-dashboard
```

### Prerequisites

* Go 1.24+
* Git
* WSL (recommended on Windows)

Verify your Go installation:

```bash
go version
```

Expected output:

```text
go version go1.24.x linux/amd64
```

### Install Dependencies

```bash
go mod tidy
```

### Build the Project

Build all packages:

```bash
go build ./...
```

Build the exporter binary:

```bash
mkdir -p bin
go build -o bin/github-team-exporter ./cmd/github-team-exporter
```

### Run the Exporter

Using Go:

```bash
go run ./cmd/github-team-exporter
```

Or using the compiled binary:

```bash
./bin/github-team-exporter
```

### Verify

Health endpoint:

```bash
curl http://localhost:2112/health
```

Expected:

```text
OK
```

Metrics endpoint:

```bash
curl http://localhost:2112/metrics
```

Expected:

```text
# HELP go_goroutines Number of goroutines
...
```

## Screenshots

### Milestone 1 – Health Endpoint

Screenshot: [docs/screenshots/milestone-1/health-endpoint.png](docs/screenshots/milestone-1/health-endpoint.png)

### Milestone 1 – Metrics Endpoint

Screenshot: [docs/screenshots/milestone-1/prometheus-metrics.png](docs/screenshots/milestone-1/prometheus-metrics.png)

### Milestone 2 – GitHub Pull Request Metrics

Screenshot: [docs/screenshots/milestone-2/github-pull-request-metrics.png](docs/screenshots/milestone-2/github-pull-request-metrics.png)

### Milestone 3 – Prometheus Target Status

Screenshot: [docs/screenshots/milestone-3/prometheus-target-up.png](docs/screenshots/milestone-3/prometheus-target-up.png)

### Milestone 3 – Pull Request Metrics Query

Screenshot: [docs/screenshots/milestone-3/prometheus-github-pull-requests-query.png](docs/screenshots/milestone-3/prometheus-github-pull-requests-query.png)

### Milestone 4 – Engineering Observability Dashboard

Screenshot: [docs/screenshots/milestone-4/grafana-dashboard-overview.png](docs/screenshots/milestone-4/grafana-dashboard-overview.png)

---

## Roadmap

### Milestone 5

- Time to first review
- Average time to merge
- Review turnaround metrics
- Review backlog metrics

### Milestone 6

- Multi-repository support
- Historical trend analysis
- Dashboard improvements
- Team-level aggregation

## Project Journey

Development progress, lessons learned, and implementation decisions are documented in: [JOURNEY.md](JOURNEY.md).

## Disclaimer

This project is a learning exercise in engineering observability, Go development, and engineering flow metrics. The design and implementation will evolve as the project progresses.
