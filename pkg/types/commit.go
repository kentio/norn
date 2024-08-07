package types

import (
	"context"
)

type GetCommitOption struct {
	Repo string
	SHA  string
}

type CreateCommitOption struct {
	Repo        string
	Tree        Tree
	SHA         string
	PickMessage string
	Target      string
	Parents     []string
}

type CheckConflictMode int

const (
	WithCommand CheckConflictMode = iota
	WithAPI
)

type CheckConflictOption struct {
	Repo     string
	Commit   string
	Target   string
	RepoPath string // only used for GitHub, because GitHub not support api for check conflict
	Mode     CheckConflictMode
	Pr       int
}

type Commit interface {
	SHA() string
	Tree() Tree
	Message() string
}

type Tree interface {
	SHA() string
	Entries() []TreeEntry
	Truncated() bool
}

type TreeEntry interface {
	SHA() string
	Path() string
	Mode() string
	Type() string
	Size() int
	Content() string
	Url() string
}

type CommitService interface {
	Get(ctx context.Context, opt *GetCommitOption) (Commit, error)
	Create(ctx context.Context, opt *CreateCommitOption) (Commit, error)
	CheckConflict(ctx context.Context, opt *CheckConflictOption) error
}
