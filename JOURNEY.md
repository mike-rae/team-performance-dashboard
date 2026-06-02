# Journey

This document records the development journey, lessons learned, challenges encountered, and decisions made while building the Engineering Observability Dashboard.

---

# 2026-06-01

## Milestone 1 – Foundations

### Objective

Create a basic Go exporter that will act as the foundation for future GitHub metrics collection and observability capabilities.

### What Was Completed

#### Repository Setup

* Created GitHub repository
* Defined project structure
* Created documentation skeleton
* Established development workflow

#### Development Environment

* Configured WSL development environment
* Installed and upgraded Go 1.24
* Configured VS Code for Go development
* Verified Go toolchain and module support

#### Source Control

* Configured GitHub SSH authentication
* Resolved repository access issues
* Connected local repository to GitHub

#### Exporter Service

Implemented a minimal Go service exposing:

```text
/health
/metrics
```

#### Prometheus Integration

Added Prometheus client library support and exposed a Prometheus-compatible metrics endpoint.

### Architecture Decision

The project will follow a lightweight exporter pattern:

```text
GitHub GraphQL API
        ↓
Go Exporter
        ↓
Prometheus
        ↓
Grafana
```

The exporter will remain stateless and rely on Prometheus for metrics storage.

### Challenges Encountered

#### VS Code Go Version Mismatch

The Go extension required a newer version than was installed within WSL.

Resolution:

* Upgraded Go to 1.24.x
* Updated module configuration
* Reloaded language services

#### GitHub Authentication Issues

Initial repository pushes failed due to authentication and permissions issues.

Resolution:

* Generated SSH key pair
* Added public key to GitHub
* Migrated repository remote from HTTPS to SSH

#### Empty Package Build Errors

Placeholder package files caused compilation failures because Go requires all source files to contain valid package declarations.

Resolution:

* Added package declarations to all internal packages

### Verification

#### Health Endpoint

Command:

```bash
curl http://localhost:2112/health
```

Result:

```text
OK
```

Evidence:

Screenshot: [docs/screenshots/milestone-1/health-endpoint.png](docs/screenshots/milestone-1/health-endpoint.png)

#### Metrics Endpoint

Command:

```bash
curl http://localhost:2112/metrics
```

Result:

```text
Prometheus metrics exposed successfully
```

Evidence:

Screenshot: [docs/screenshots/milestone-1/prometheus-metrics.png](docs/screenshots/milestone-1/prometheus-metrics.png)

### Lessons Learned

* WSL provides a robust Go development environment once configured correctly.
* Establishing package boundaries early simplifies future development.
* Prometheus integration in Go is straightforward.
* SSH authentication removes many common GitHub credential issues.
* Small, incremental milestones reduce project risk and improve learning.

### Next Milestone

* Implement configuration loading
* Create GitHub GraphQL client
* Execute first authenticated GitHub query
* Retrieve pull request data
* Expose first GitHub-derived metric

### Status

Milestone 1 Complete ✅

---

# 2026-06-02

## Milestone 2 – GitHub GraphQL Integration

### Objective

Connect the Go exporter to GitHub using the GitHub GraphQL API and expose the first GitHub-derived Prometheus metrics.

### What Was Completed

* Added environment configuration loading from `.env`
* Added GitHub token, owner, and repository configuration
* Created GitHub GraphQL client
* Authenticated with GitHub using a personal access token
* Queried pull request counts by state
* Refactored pull request queries into a reusable `PullRequestCount` function
* Added custom Prometheus metric for pull requests
* Exposed pull request counts using labels:

  * `state="open"`
  * `state="closed"`
  * `state="merged"`

### Current Metric

```text
github_pull_requests{owner="mike-rae",repo="engineering-observability-dashboard",state="open"} 0
github_pull_requests{owner="mike-rae",repo="engineering-observability-dashboard",state="closed"} 1
github_pull_requests{owner="mike-rae",repo="engineering-observability-dashboard",state="merged"} 1
```

### Verification

Command:

```bash
curl http://localhost:2112/metrics | grep github_pull_requests
```

Screenshot: [docs/screenshots/milestone-2/github-pull-request-metrics.png](docs/screenshots/milestone-2/github-pull-request-metrics.png)

### Challenges Encountered

#### GraphQL State Argument

The first version passed a single pull request state directly to the GraphQL query.

GitHub expected a list of pull request states, so the query needed to use:

```go
graphql:"pullRequests(states: [$state])"
```

### Lessons Learned

* GitHub GraphQL is a good fit for targeted engineering metrics.
* Prometheus labels are better than separate metric names for related states.
* A metric returning `0` is still a valid and useful result.
* Small reusable query functions keep the exporter simpler as more metrics are added.

### Next Milestone

* Add Docker Compose
* Run the exporter in a container
* Add Prometheus
* Configure Prometheus to scrape the exporter
* Prepare for Grafana dashboarding

### Status

Milestone 2 Complete ✅
