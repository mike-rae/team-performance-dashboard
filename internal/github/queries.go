package github

import (
	"context"
	"time"

	"github.com/shurcooL/githubv4"
)

type PullRequestDetails struct {
	Number              int
	Title               string
	CreatedAt           time.Time
	ReviewRequestsCount int
	ReviewsCount        int
}

func PullRequestCount(client *githubv4.Client, owner string, repo string, state string) (int, error) {
	var query struct {
		Repository struct {
			PullRequests struct {
				TotalCount int
			} `graphql:"pullRequests(states: [$state])"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}

	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"repo":  githubv4.String(repo),
		"state": githubv4.PullRequestState(state),
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return 0, err
	}

	return query.Repository.PullRequests.TotalCount, nil
}

func OpenPullRequestDetails(client *githubv4.Client, owner string, repo string) ([]PullRequestDetails, error) {
	var query struct {
		Repository struct {
			PullRequests struct {
				Nodes []struct {
					Number    int
					Title     string
					CreatedAt githubv4.DateTime

					ReviewRequests struct {
						TotalCount int
					} `graphql:"reviewRequests(first: 10)"`

					Reviews struct {
						TotalCount int
					} `graphql:"reviews(first: 10)"`
				}
			} `graphql:"pullRequests(first: 50, states: OPEN, orderBy: {field: CREATED_AT, direction: DESC})"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}

	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"repo":  githubv4.String(repo),
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}

	results := make([]PullRequestDetails, 0, len(query.Repository.PullRequests.Nodes))

	for _, pr := range query.Repository.PullRequests.Nodes {
		results = append(results, PullRequestDetails{
			Number:              pr.Number,
			Title:               pr.Title,
			CreatedAt:           pr.CreatedAt.Time,
			ReviewRequestsCount: pr.ReviewRequests.TotalCount,
			ReviewsCount:        pr.Reviews.TotalCount,
		})
	}

	return results, nil
}
