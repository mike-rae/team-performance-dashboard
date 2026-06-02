# Engineering Observability Dashboard

A build-it-yourself engineering observability platform built with Go, GitHub GraphQL, Prometheus, Grafana, and Docker.

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

* Repository structure established
* Go module initialised
* Local development environment configured
* Health endpoint implemented
* Metrics endpoint implemented

### Milestone 2 – GitHub GraphQL Integration Complete ✅

Completed:

Environment configuration loading
GitHub GraphQL client
GitHub authentication using personal access token
Pull request count query
Custom Prometheus metric
Pull request metrics by state

### Current Functionality

The exporter currently provides:

* `/health`
* `/metrics`
* GitHub pull request metrics by state

Current custom metric:

```bash
github_pull_requests{owner="mike-rae",repo="engineering-observability-dashboard",state="open"}
github_pull_requests{owner="mike-rae",repo="engineering-observability-dashboard",state="closed"}
github_pull_requests{owner="mike-rae",repo="engineering-observability-dashboard",state="merged"}
```

The project does not yet:

* Run through Docker Compose
* Configure Prometheus scraping
* Provide Grafana dashboards

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

## Roadmap

### Milestone 2

* Environment configuration
* GitHub GraphQL client
* First authenticated GitHub query
* Pull request collection

### Milestone 3

* Custom GitHub metrics
* Prometheus scrape configuration
* Docker Compose setup

### Milestone 4

* Grafana integration
* Dashboard provisioning
* First engineering observability dashboard

### Milestone 5

* Multi-repository support
* Additional engineering flow metrics
* Trend analysis

## Project Journey

Development progress, lessons learned, and implementation decisions are documented in: [JOURNEY.md](JOURNEY.md).

## Disclaimer

This project is a learning exercise in engineering observability, Go development, and engineering flow metrics. The design and implementation will evolve as the project progresses.
