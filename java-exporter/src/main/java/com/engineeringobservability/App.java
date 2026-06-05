package com.engineeringobservability;

import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;

import com.sun.net.httpserver.HttpServer;

import io.prometheus.metrics.core.metrics.Gauge;
import io.prometheus.metrics.exporter.httpserver.HTTPServer;

public class App {
    public static void main(String[] args) throws IOException {
        startHealthServer();

        Gauge.builder()
                .name("java_exporter_metadata")
                .help("Java exporter information.")
                .register()
                .set(1);

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
