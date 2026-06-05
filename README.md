# Engineering Observability Dashboard

A lightweight engineering observability platform built using Go, GitHub GraphQL, Prometheus and Grafana.

## Dashboard Preview

![Engineering Observability Dashboard](docs/screenshots/milestone-5/grafana-dashboard-overview.png)

## Overview

This project explores how engineering flow metrics can be collected from GitHub and visualised using open-source observability tooling.

The platform collects pull request data from GitHub using GraphQL, exposes custom Prometheus metrics through a Go exporter, stores metrics in Prometheus and visualises engineering insights in Grafana.

## Architecture

```text
GitHub GraphQL API
        ↓
Go Exporter
        ↓
Prometheus
        ↓
Grafana
```

## Current Metrics

### Pull Request Volume

* Open Pull Requests
* Closed Pull Requests
* Merged Pull Requests

### Pull Request Health

* Stale Pull Requests
* Oldest Open Pull Request
* Average Open Pull Request Age

### Review Flow

* Review Backlog
* Waiting For Review
* Average Time To First Review

### Delivery Flow

* Average Time To Merge

## Current Status

| Milestone                   | Status |
| --------------------------- | ------ |
| M1 Foundations              | ✅      |
| M2 GitHub GraphQL           | ✅      |
| M3 Docker & Prometheus      | ✅      |
| M4 Grafana Dashboard        | ✅      |
| M5 Engineering Flow Metrics | ✅      |
| M6 Dashboard Provisioning   | ✅      |

## Quick Start

```bash
git clone git@github.com:mike-rae/engineering-observability-dashboard.git
cd engineering-observability-dashboard

docker compose up --build -d
```

Open:

* Grafana: http://localhost:3000
* Prometheus: http://localhost:9090
* Metrics: http://localhost:2112/metrics

## Documentation

* [Setup Guide](docs/SETUP.md)
* [Project Journey](JOURNEY.md)

## Roadmap

### Milestone 7

* Multi-repository support
* Team aggregation
* Historical trend analysis

### Milestone 8

* Java implementation
* Go vs Java comparison
* Performance comparison
* Maintainability comparison

## Disclaimer

This project is a learning exercise in engineering observability, Go development and engineering flow metrics.
