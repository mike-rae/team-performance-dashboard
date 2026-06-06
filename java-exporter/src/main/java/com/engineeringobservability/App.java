package com.engineeringobservability;

import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;
import java.time.Duration;
import java.time.Instant;
import java.util.List;

import com.engineeringobservability.config.AppConfig;
import com.engineeringobservability.github.GitHubClient;
import com.engineeringobservability.github.PullRequestDetails;
import com.engineeringobservability.metrics.MetricsExporter;
import com.sun.net.httpserver.HttpServer;

import io.prometheus.metrics.exporter.httpserver.HTTPServer;

public class App {

    public static void main(String[] args) throws IOException, InterruptedException {
        AppConfig config = AppConfig.load();

        System.out.printf(
                "Loaded config for %s/%s%n",
                config.githubOwner(),
                config.githubRepo());

        GitHubClient gitHubClient = new GitHubClient(config.githubToken());

        // get PR states
        String[] states = { "OPEN", "CLOSED", "MERGED" };
        for (String state : states) {

            int count = gitHubClient.pullRequestCount(
                    config.githubOwner(),
                    config.githubRepo(),
                    state);

            MetricsExporter.pullRequests
                    .labelValues(
                            config.githubOwner(),
                            config.githubRepo(),
                            state.toLowerCase())
                    .set(count);

            System.out.printf(
                    "%s PRs: %d%n",
                    state,
                    count);
        }

        // get pull request ages for open PRs
        List<PullRequestDetails> openPullRequests = gitHubClient.openPullRequestDetails(
                config.githubOwner(),
                config.githubRepo());

        int staleThresholdDays = 2;
        int staleCount = 0;

        for (PullRequestDetails pr : openPullRequests) {
            double ageDays = Duration.between(pr.createdAt(), Instant.now()).toHours() / 24.0;

            MetricsExporter.pullRequestAgeDays
                    .labelValues(
                            config.githubOwner(),
                            config.githubRepo(),
                            String.valueOf(pr.number()))
                    .set(ageDays);

            if (ageDays >= staleThresholdDays) {
                staleCount++;
            }
        }

        MetricsExporter.pullRequestsStale
                .labelValues(
                        config.githubOwner(),
                        config.githubRepo())
                .set(staleCount);

        startHealthServer();

        HTTPServer metricsServer = HTTPServer.builder()
                .port(2113)
                .buildAndStart();

        System.out.println("java-exporter running");
        System.out.println("health endpoint:  http://localhost:8081/health");
        System.out.println("metrics endpoint: http://localhost:2113/metrics");

        Runtime.getRuntime().addShutdownHook(new Thread(metricsServer::close));
    }

    private static void startHealthServer() throws IOException {
        HttpServer server = HttpServer.create(new InetSocketAddress(8081), 0);

        server.createContext("/health", exchange -> {
            String response = "OK\n";
            exchange.sendResponseHeaders(200, response.length());

            try (OutputStream os = exchange.getResponseBody()) {
                os.write(response.getBytes());
            }
        });

        server.start();
    }
}
