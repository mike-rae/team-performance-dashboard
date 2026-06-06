package com.engineeringobservability.github;

import java.time.Instant;

public record PullRequestDetails(
        int number,
        Instant createdAt) {
}
