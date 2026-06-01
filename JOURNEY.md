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
