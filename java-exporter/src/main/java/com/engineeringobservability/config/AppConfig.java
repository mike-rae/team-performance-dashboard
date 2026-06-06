package com.engineeringobservability.config;

import io.github.cdimascio.dotenv.Dotenv;

public record AppConfig(
        String githubToken,
        String githubOwner,
        String githubRepo) {
    public static AppConfig load() {
        Dotenv dotenv = Dotenv.configure()
                .directory("../")
                .ignoreIfMissing()
                .load();

        String token = requireEnv("GITHUB_TOKEN", dotenv);
        String owner = requireEnv("GITHUB_OWNER", dotenv);
        String repo = requireEnv("GITHUB_REPO", dotenv);

        return new AppConfig(token, owner, repo);
    }

    private static String requireEnv(String name, Dotenv dotenv) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            value = dotenv.get(name);
        }

        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " is required");
        }

        return value;
    }
}
