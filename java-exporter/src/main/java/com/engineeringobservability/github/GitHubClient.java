package com.engineeringobservability.github;

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.time.Instant;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

public class GitHubClient {
  private static final String GITHUB_GRAPHQL_URL = "https://api.github.com/graphql";

  private final String token;
  private final HttpClient httpClient;
  private final ObjectMapper objectMapper;

  public GitHubClient(String token) {
    this.token = token;
    this.httpClient = HttpClient.newHttpClient();
    this.objectMapper = new ObjectMapper();
  }

  public int pullRequestCount(String owner, String repo, String state) throws IOException, InterruptedException {
    String query = """
        query($owner: String!, $repo: String!, $state: [PullRequestState!]) {
          repository(owner: $owner, name: $repo) {
            pullRequests(states: $state) {
              totalCount
            }
          }
        }
        """;

    Map<String, Object> body = Map.of(
        "query", query,
        "variables", Map.of(
            "owner", owner,
            "repo", repo,
            "state", java.util.List.of(state)));

    String requestBody = objectMapper.writeValueAsString(body);

    HttpRequest request = HttpRequest.newBuilder()
        .uri(URI.create(GITHUB_GRAPHQL_URL))
        .header("Authorization", "Bearer " + token)
        .header("Content-Type", "application/json")
        .POST(HttpRequest.BodyPublishers.ofString(requestBody))
        .build();

    HttpResponse<String> response = httpClient.send(
        request,
        HttpResponse.BodyHandlers.ofString());

    if (response.statusCode() != 200) {
      throw new IOException("GitHub GraphQL request failed: " + response.statusCode() + " " + response.body());
    }

    JsonNode root = objectMapper.readTree(response.body());

    JsonNode errors = root.get("errors");
    if (errors != null && errors.size() > 0) {
      throw new IOException("GitHub GraphQL errors: " + errors);
    }

    return root
        .path("data")
        .path("repository")
        .path("pullRequests")
        .path("totalCount")
        .asInt();
  }

  public List<PullRequestDetails> openPullRequestDetails(
      String owner,
      String repo) throws IOException, InterruptedException {
    String query = """
        query($owner: String!, $repo: String!) {
          repository(owner: $owner, name: $repo) {
            pullRequests(first: 50, states: OPEN, orderBy: {field: CREATED_AT, direction: DESC}) {
              nodes {
                number
                createdAt
              }
            }
          }
        }
        """;

    Map<String, Object> body = Map.of(
        "query", query,
        "variables", Map.of(
            "owner", owner,
            "repo", repo));

    String requestBody = objectMapper.writeValueAsString(body);

    HttpRequest request = HttpRequest.newBuilder()
        .uri(URI.create(GITHUB_GRAPHQL_URL))
        .header("Authorization", "Bearer " + token)
        .header("Content-Type", "application/json")
        .POST(HttpRequest.BodyPublishers.ofString(requestBody))
        .build();

    HttpResponse<String> response = httpClient.send(
        request,
        HttpResponse.BodyHandlers.ofString());

    if (response.statusCode() != 200) {
      throw new IOException("GitHub GraphQL request failed: " + response.statusCode() + " " + response.body());
    }

    JsonNode root = objectMapper.readTree(response.body());

    JsonNode errors = root.get("errors");
    if (errors != null && errors.size() > 0) {
      throw new IOException("GitHub GraphQL errors: " + errors);
    }

    JsonNode nodes = root
        .path("data")
        .path("repository")
        .path("pullRequests")
        .path("nodes");

    List<PullRequestDetails> results = new ArrayList<>();

    for (JsonNode node : nodes) {
      results.add(new PullRequestDetails(
          node.path("number").asInt(),
          Instant.parse(node.path("createdAt").asText())));
    }

    return results;
  }
}
