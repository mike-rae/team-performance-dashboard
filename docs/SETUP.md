# Setup Guide

## Clone the Repository

```bash
git clone git@github.com:mike-rae/engineering-observability-dashboard.git
cd engineering-observability-dashboard
```

## Prerequisites

* Go 1.24+
* Docker
* Docker Compose
* Git

## Build

```bash
go build ./...
```

## Run Locally

```bash
go run ./cmd/github-team-exporter
```

## Run with Docker

```bash
docker compose up --build -d
```

## Verify

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

Prometheus:

```text
http://localhost:9090
```

Grafana:

```text
http://localhost:3000
```

## Screenshots

### Milestone 1

* docs/screenshots/milestone-1/health-endpoint.png
* docs/screenshots/milestone-1/prometheus-metrics.png

### Milestone 2

* docs/screenshots/milestone-2/github-pull-request-metrics.png

### Milestone 3

* docs/screenshots/milestone-3/prometheus-target-up.png
* docs/screenshots/milestone-3/prometheus-github-pull-requests-query.png

### Milestone 4

* docs/screenshots/milestone-4/grafana-dashboard-overview.png

### Milestone 5

* docs/screenshots/milestone-5/grafana-dashboard-overview.png
