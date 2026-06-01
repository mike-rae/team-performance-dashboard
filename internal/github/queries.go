package github

import (
	"context"

	"github.com/shurcooL/githubv4"
)

func OpenPullRequestCount(client *githubv4.Client, owner string, repo string) (int, error) {
	var query struct {
		Repository struct {
			PullRequests struct {
				TotalCount int
			} `graphql:"pullRequests(states: OPEN)"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}

	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"repo":  githubv4.String(repo),
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return 0, err
	}

	return query.Repository.PullRequests.TotalCount, nil
}
