package com.engineeringobservability.config;

public record AppConfig(
        String githubToken,
        String githubOwner,
        String githubRepo) {
    public static AppConfig load() {
        String token = requireEnv("GITHUB_TOKEN");
        String owner = requireEnv("GITHUB_OWNER");
        String repo = requireEnv("GITHUB_REPO");

        return new AppConfig(token, owner, repo);
    }

    private static String requireEnv(String name) {
        String value = System.getenv(name);

        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " is required");
        }

        return value;
    }
}
