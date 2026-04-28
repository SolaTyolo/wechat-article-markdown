package storage

import (
	"context"

	"github.com/SolaTyolo/wechat-article-markdown/pkg/wechatmd/core"
)

// Storage persists article markdown and image assets.
type Storage interface {
	Save(ctx context.Context, article *core.ArticleResult) error
}
