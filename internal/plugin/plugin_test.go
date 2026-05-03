package plugin_test

import (
	"net/http"
	"testing"

	_ "github.com/kbsink-org/kbsink/internal/plugin"
	"github.com/kbsink-org/kbsink/pkg/pluginreg"
)

func TestBuiltinPlugins_registered(t *testing.T) {
	for _, name := range []string{"wechat", "WeChat", "xhs", "XHS"} {
		pl, ok := pluginreg.Lookup(name)
		if !ok || pl == nil {
			t.Fatalf("Lookup(%q) missing", name)
		}
		p, d, err := pl.NewComponents(http.DefaultClient)
		if err != nil {
			t.Fatal(err)
		}
		if p == nil {
			t.Fatalf("parser nil for %q", name)
		}
		if d == nil {
			t.Fatalf("builtin %q should include a driver", name)
		}
	}
}
