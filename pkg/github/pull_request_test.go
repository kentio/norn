package github

import (
	"context"
	"github.com/kentio/norn/pkg/types"
	"testing"
)

func TestPullRequestService_Get(t *testing.T) {

	ctx := context.Background()
	token := ""
	client := NewGithubClient(ctx, token)

	pullRequestClient := NewPullRequestService(client)

	pr, err := pullRequestClient.Get(ctx, &types.GetMergeRequestOption{
		Repo:    "",
		MergeID: "",
	})
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	t.Logf("pr: %+v state %s", pr, pr.State().String())

}
