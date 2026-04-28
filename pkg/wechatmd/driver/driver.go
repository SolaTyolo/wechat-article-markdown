package driver

import (
	"context"

	"github.com/solat/wechat-article-markdown/pkg/wechatmd/core"
)

// Driver fetches article page source by URL.
type Driver interface {
	Fetch(ctx context.Context, url string) (*core.FetchResult, error)
}
