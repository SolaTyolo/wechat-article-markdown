package driver

import (
	"net/http"
	"strings"

	"github.com/kbsink-org/kbsink/internal/htmldriver"
	"github.com/kbsink-org/kbsink/pkg/core"
)

const (
	defaultUserAgent = "Mozilla/5.0 (compatible; wechatmd/1.0)"
	// wechatUserAgent matches default historical behaviour for mp.weixin.qq.com pages.
	wechatUserAgent = defaultUserAgent
	// xhsUserAgent is a common mobile Safari UA; XHS pages may vary by client hints.
	xhsUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/147.0.0.0 Safari/537.36"
)

// NewHTMLDriver returns an HTTP HTML fetch driver. If userAgent is empty or whitespace, defaultUserAgent is used.
func NewHTMLDriver(client *http.Client, userAgent string) core.Driver {
	ua := strings.TrimSpace(userAgent)
	if ua == "" {
		ua = defaultUserAgent
	}
	return htmldriver.New(client, ua)
}

// NewWechatDriver returns the HTML fetch driver paired with WeChat article parsing.
func NewWechatDriver(client *http.Client) core.Driver {
	return htmldriver.New(client, wechatUserAgent)
}

// NewXHSDriver returns the HTML fetch driver paired with XHS (小红书) parsing.
func NewXHSDriver(client *http.Client) core.Driver {
	return htmldriver.New(client, xhsUserAgent)
}

// DefaultXHSUserAgent is the User-Agent used by NewXHSDriver; reuse for CDN image requests.
func DefaultXHSUserAgent() string { return xhsUserAgent }
