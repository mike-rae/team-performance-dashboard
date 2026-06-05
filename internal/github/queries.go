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
	FirstReviewAt       *time.Time
	TimeToMergeHours    float64
}

type MergedPullRequest struct {
	Number    int
	CreatedAt time.Time
	MergedAt  time.Time
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
						Nodes      []struct {
							CreatedAt githubv4.DateTime
						}
					} `graphql:"reviews(first: 1)"`
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
		var firstReviewAt *time.Time

		if len(pr.Reviews.Nodes) > 0 {
			t := pr.Reviews.Nodes[0].CreatedAt.Time
			firstReviewAt = &t
		}

		results = append(results, PullRequestDetails{
			Number:              pr.Number,
			Title:               pr.Title,
			CreatedAt:           pr.CreatedAt.Time,
			ReviewRequestsCount: pr.ReviewRequests.TotalCount,
			ReviewsCount:        pr.Reviews.TotalCount,
			FirstReviewAt:       firstReviewAt,
		})
	}

	return results, nil
}

func MergedPullRequests(
	client *githubv4.Client,
	owner string,
	repo string,
) ([]MergedPullRequest, error) {

	var query struct {
		Repository struct {
			PullRequests struct {
				Nodes []struct {
					Number    int
					CreatedAt githubv4.DateTime
					MergedAt  githubv4.DateTime
				}
			} `graphql:"pullRequests(first: 50, states: MERGED, orderBy: {field: UPDATED_AT, direction: DESC})"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}

	vars := map[string]interface{}{
		"owner": githubv4.String(owner),
		"repo":  githubv4.String(repo),
	}

	err := client.Query(context.Background(), &query, vars)
	if err != nil {
		return nil, err
	}

	results := make([]MergedPullRequest, 0)

	for _, pr := range query.Repository.PullRequests.Nodes {

		results = append(results, MergedPullRequest{
			Number:    pr.Number,
			CreatedAt: pr.CreatedAt.Time,
			MergedAt:  pr.MergedAt.Time,
		})
	}

	return results, nil
}
