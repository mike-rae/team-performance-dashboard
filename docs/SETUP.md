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

---

## Java Exporter

The Java exporter is a technology spike used to compare Java and Go implementations of a GitHub Prometheus exporter.

### Development Mode

Run the exporter directly from your IDE or Maven.

```bash
cd java-exporter

mvn exec:java \
  -Dexec.mainClass="com.engineeringobservability.App"
```

Verify:

```bash
curl http://localhost:8081/health
```

Expected:

```text
OK
```

Metrics:

```bash
curl -s http://localhost:2113/metrics | grep java_github
```

### Docker Mode

Build and run the Java exporter container:

```bash
docker compose up --build -d
```

Verify:

```bash
docker compose ps
```

Expected:

```text
java-exporter    Up
```

Health:

```bash
curl http://localhost:8081/health
```

Metrics:

```bash
curl http://localhost:2113/metrics
```

### Prometheus Integration

#### Development Mode

If Prometheus runs in Docker and the Java exporter runs locally:

```yaml
- job_name: java-github-exporter
  static_configs:
    - targets:
        - host.docker.internal:2113
```

Linux and WSL users may also require:

```yaml
extra_hosts:
  - "host.docker.internal:host-gateway"
```

on the Prometheus service.

#### Docker Mode

If both Prometheus and the Java exporter run inside Docker Compose:

```yaml
- job_name: java-github-exporter
  static_configs:
    - targets:
        - java-exporter:2113
```

### Implemented Metrics

* java_github_pull_requests
* java_github_pull_request_age_days
* java_github_pull_requests_stale

---

## Screenshots

### Milestone 1

* [health-endpoint.png](docs/screenshots/milestone-1/health-endpoint.png)
* [prometheus-metrics.png](docs/screenshots/milestone-1/prometheus-metrics.png)

### Milestone 2

* [github-pull-request-metrics.png](docs/screenshots/milestone-2/github-pull-request-metrics.png)

### Milestone 3

* [prometheus-target-up.png](docs/screenshots/milestone-3/prometheus-target-up.png)
* [prometheus-github-pull-requests-query.png](docs/screenshots/milestone-3/prometheus-github-pull-requests-query.png)

### Milestone 4

* [grafana-dashboard-overview.png](docs/screenshots/milestone-4/grafana-dashboard-overview.png)

### Milestone 5

* [grafana-dashboard-overview.png](docs/screenshots/milestone-5/grafana-dashboard-overview.png)

### Milestone 7

* [java-exporter-metrics.png](docs/screenshots/milestone-7/java-exporter-metrics.png)
* [java-prometheus-target-up.png](docs/screenshots/milestone-7/java-prometheus-target-up.png)
