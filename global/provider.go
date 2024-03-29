package global

import (
	"context"
	"github.com/kentio/norn/github"
	"github.com/kentio/norn/types"
	"github.com/sirupsen/logrus"
)

// NewProvider NewClient returns a new client for the given vendor.
func NewProvider(ctx context.Context, vendor string, token string) (types.Provider, error) {
	logrus.Debugf("New provider: %s", vendor)

	switch vendor {
	case "gh", "github":
		return github.NewProvider(ctx, token), nil
	default:
		return nil, types.ErrUnknownProvider
	}
}
