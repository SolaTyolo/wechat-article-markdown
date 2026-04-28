package driver

import (
	"context"

	"github.com/SolaTyolo/wechat-article-markdown/pkg/wechatmd/core"
)

// Driver fetches article page source by URL.
type Driver interface {
	Fetch(ctx context.Context, url string) (*core.FetchResult, error)
}
