package driver

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/SolaTyolo/wechat-article-markdown/pkg/wechatmd/core"
)

// HTMLDriver fetches HTML from a page URL through plain HTTP.
type HTMLDriver struct {
	client *http.Client
}

func NewHTMLDriver(client *http.Client) *HTMLDriver {
	if client == nil {
		client = http.DefaultClient
	}
	return &HTMLDriver{client: client}
}

func (d *HTMLDriver) Fetch(ctx context.Context, rawURL string) (*core.FetchResult, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; wechatmd/1.0)")

	resp, err := d.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &core.FetchResult{
		URL:  resp.Request.URL.String(),
		HTML: string(body),
	}, nil
}
