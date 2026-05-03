package wechat

import (
	"net/http"

	"github.com/kbsink-org/kbsink/internal/articleparse"
	"github.com/kbsink-org/kbsink/pkg/core"
	"github.com/kbsink-org/kbsink/pkg/driver"
)

// Plugin is the built-in WeChat article plugin (parser + fetch driver).
type Plugin struct{}

// New returns a Plugin registered as "wechat".
func New() core.Plugin {
	return Plugin{}
}

func (Plugin) Name() string { return "wechat" }

func (Plugin) NewComponents(client *http.Client) (core.Parser, core.Driver, error) {
	return articleparse.NewWechatParser(), driver.NewWechatDriver(client), nil
}
