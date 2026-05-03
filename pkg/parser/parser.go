package parser

import (
	"github.com/kbsink-org/kbsink/internal/articleparse"
	"github.com/kbsink-org/kbsink/pkg/core"
)

// NewWechatParser returns the built-in WeChat article parser.
func NewWechatParser() core.Parser {
	return articleparse.NewWechatParser()
}

// NewXHSParser returns the built-in XHS (小红书) HTML parser.
func NewXHSParser() core.Parser {
	return articleparse.NewXiaohongshuParser()
}
