package common

import (
	"context"
	"github.com/kentio/norn/pkg/github"
	tp "github.com/kentio/norn/pkg/types"
	"github.com/sirupsen/logrus"
)

// NewProvider NewClient returns a new client for the given vendor.
func NewProvider(ctx context.Context, vendor string, opt *tp.CreateProviderOption) (tp.Provider, error) {
	logrus.Debugf("New provider: %s", vendor)

	switch vendor {
	case "gh", "github":
		return github.NewProvider(ctx, opt), nil
	default:
		return nil, tp.ErrUnknownProvider
	}
}
