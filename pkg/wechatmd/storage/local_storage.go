package storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/solat/wechat-article-markdown/pkg/wechatmd/core"
)

const defaultOutputRoot = "output"

// LocalStorage saves markdown and images into local filesystem.
type LocalStorage struct {
	root string
}

func NewLocalStorage(root string) *LocalStorage {
	if root == "" {
		root = defaultOutputRoot
	}
	return &LocalStorage{root: root}
}

func (s *LocalStorage) Save(_ context.Context, article *core.ArticleResult) error {
	if article == nil {
		return fmt.Errorf("article is nil")
	}
	baseDir := filepath.FromSlash(article.OutputDir)
	if err := os.MkdirAll(baseDir, 0o755); err != nil {
		return err
	}
	imageDir := filepath.Join(baseDir, "images")
	if err := os.MkdirAll(imageDir, 0o755); err != nil {
		return err
	}

	for _, img := range article.Images {
		target := filepath.Join(baseDir, filepath.FromSlash(img.RelativePath))
		if err := os.WriteFile(target, img.Data, 0o644); err != nil {
			return err
		}
	}
	mdPath := filepath.FromSlash(article.MarkdownPath)
	if err := os.WriteFile(mdPath, []byte(article.Markdown), 0o644); err != nil {
		return err
	}
	return nil
}
