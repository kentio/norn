package github

import (
	"context"
	gh "github.com/google/go-github/v62/github"
	tp "github.com/kentio/norn/pkg/types"
)

type Provider struct {
	providerID tp.ProviderType
	client     *gh.Client

	commitService       *CommitService
	referenceService    *ReferenceService
	mergeRequestService *PullRequestService
	commentService      *CommentService
	pickService         *PickService
	repositoryService   *RepositoryService
}

func NewProvider(ctx context.Context, opt *tp.CreateProviderOption) *Provider {
	var client *gh.Client
	if opt.BaseUrl != nil {
		client = NewGitHubWithBaseUrl(ctx, opt)
	} else {
		client = NewGithubClient(ctx, opt.Token)
	}
	return &Provider{
		providerID:          tp.GitHubProvider,
		commitService:       NewCommitService(client),
		referenceService:    NewReferenceService(client),
		mergeRequestService: NewPullRequestService(client),
		commentService:      NewCommentService(client),
		pickService:         NewPickService(client),
		repositoryService:   NewRepositoryService(client),
	}
}

// NewProviderWithClient creates a new provider with the given client.
func NewProviderWithClient(client *gh.Client) *Provider {
	return &Provider{
		providerID:          tp.GitHubProvider,
		client:              client,
		commitService:       NewCommitService(client),
		referenceService:    NewReferenceService(client),
		mergeRequestService: NewPullRequestService(client),
		commentService:      NewCommentService(client),
		pickService:         NewPickService(client),
		repositoryService:   NewRepositoryService(client),
	}
}

func (p *Provider) Commit() tp.CommitService {
	return p.commitService
}

func (p *Provider) Reference() tp.ReferenceService {
	return p.referenceService
}

func (p *Provider) MergeRequest() tp.MergeRequestService {
	return p.mergeRequestService
}

func (p *Provider) Comment() tp.CommentService {
	return p.commentService
}

func (p *Provider) Repository() tp.RepositoryService {
	return p.repositoryService
}

func (p *Provider) ProviderID() tp.ProviderType {
	return p.providerID
}

func (p *Provider) Pick() tp.PickService {
	return p.pickService
}
