package parser

import (
	"context"

	"github.com/solat/wechat-article-markdown/pkg/wechatmd/core"
)

// Parser transforms raw HTML into metadata + markdown + image references.
type Parser interface {
	Parse(ctx context.Context, fetched *core.FetchResult, outputDir string) (*core.ArticleResult, error)
}
