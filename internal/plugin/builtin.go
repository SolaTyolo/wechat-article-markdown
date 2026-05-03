// Package plugin registers built-in kbsink plugins (wechat, xhs) with pluginreg.
// Blank-import from the kbsink module (e.g. cmd/kb-sink-md/cli) to load built-ins.
package plugin

import (
	"github.com/kbsink-org/kbsink/internal/plugin/wechat"
	"github.com/kbsink-org/kbsink/internal/plugin/xhs"
	"github.com/kbsink-org/kbsink/pkg/pluginreg"
)

func init() {
	pluginreg.Register(wechat.New())
	pluginreg.Register(xhs.New())
}
