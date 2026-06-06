package com.engineeringobservability;

import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;

import com.engineeringobservability.config.AppConfig;
import com.engineeringobservability.github.GitHubClient;
import com.sun.net.httpserver.HttpServer;

import io.prometheus.metrics.core.metrics.Gauge;
import io.prometheus.metrics.exporter.httpserver.HTTPServer;

public class App {
    private static final Gauge exporterMetadata = Gauge.builder()
            .name("java_exporter_metadata")
            .help("Java exporter information.")
            .register();

    private static final Gauge pullRequests = Gauge.builder()
            .name("java_github_pull_requests")
            .help("Number of GitHub pull requests by state.")
            .labelNames("owner", "repo", "state")
            .register();

    public static void main(String[] args) throws IOException, InterruptedException {
        AppConfig config = AppConfig.load();

        System.out.printf(
                "Loaded config for %s/%s%n",
                config.githubOwner(),
                config.githubRepo());

        GitHubClient gitHubClient = new GitHubClient(config.githubToken());

        String[] states = { "OPEN", "CLOSED", "MERGED" };

        for (String state : states) {

            int count = gitHubClient.pullRequestCount(
                    config.githubOwner(),
                    config.githubRepo(),
                    state);

            pullRequests
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

        startHealthServer();

        exporterMetadata.set(2888);

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
