package storage

import (
	"context"

	"github.com/solat/wechat-article-markdown/pkg/wechatmd/core"
)

// Storage persists article markdown and image assets.
type Storage interface {
	Save(ctx context.Context, article *core.ArticleResult) error
}
