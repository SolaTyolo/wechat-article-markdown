package xhs

import (
	"net/http"

	"github.com/kbsink-org/kbsink/internal/articleparse"
	"github.com/kbsink-org/kbsink/pkg/core"
	"github.com/kbsink-org/kbsink/pkg/driver"
)

// Plugin is the built-in XHS (小红书) plugin: parser + fetch driver.
type Plugin struct{}

// New returns a Plugin registered as "xhs".
func New() core.Plugin {
	return Plugin{}
}

func (Plugin) Name() string { return "xhs" }

func (Plugin) NewComponents(client *http.Client) (core.Parser, core.Driver, error) {
	return articleparse.NewXiaohongshuParser(), driver.NewXHSDriver(client), nil
}
