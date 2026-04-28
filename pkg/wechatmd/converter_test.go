package wechatmd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/solat/wechat-article-markdown/pkg/wechatmd/core"
	prs "github.com/solat/wechat-article-markdown/pkg/wechatmd/parser"
)

type memoryStorage struct {
	saved *core.ArticleResult
}

func (m *memoryStorage) Save(_ context.Context, article *core.ArticleResult) error {
	m.saved = article
	return nil
}

func TestSanitizeFileName(t *testing.T) {
	got := sanitizeFileName(`  hello:/\*world?  `)
	if got != "hello____world" {
		t.Fatalf("unexpected sanitized file name: %q", got)
	}
}

func TestConvertWithDefaultHTMLDriverAndParser(t *testing.T) {
	imageData := []byte{1, 2, 3}
	mux := http.NewServeMux()
	mux.HandleFunc("/article", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`
<html>
  <head><title>fallback-title</title></head>
  <body>
    <h1 id="activity-name">Test Article</h1>
    <strong id="js_name">Test Account</strong>
    <em id="publish_time">2026-04-28 10:20:30</em>
    <div id="js_content">
      <p>Hello</p>
      <img data-src="` + "http://example.invalid/img.jpg" + `" />
    </div>
  </body>
</html>`))
	})
	mux.HandleFunc("/img.jpg", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "image/jpeg")
		_, _ = w.Write(imageData)
	})
	ts := httptest.NewServer(mux)
	defer ts.Close()

	// Ensure image src points to our test server.
	articleURL := ts.URL + "/article"
	html := `
<html><body>
<h1 id="activity-name">Test Article</h1>
<strong id="js_name">Test Account</strong>
<em id="publish_time">2026-04-28 10:20:30</em>
<div id="js_content"><p>Hello</p><img data-src="` + ts.URL + `/img.jpg" /></div>
</body></html>`
	driver := &stubDriver{res: &core.FetchResult{URL: articleURL, HTML: html}}
	memStore := &memoryStorage{}

	c := NewConverter(
		WithDriver(driver),
		WithParser(prs.NewWechatParser()),
		WithStorage(memStore),
		WithHTTPClient(ts.Client()),
	)
	res, err := c.Convert(context.Background(), articleURL, core.ConvertOptions{OutputRoot: "output"})
	if err != nil {
		t.Fatalf("convert error: %v", err)
	}

	if res.Title != "Test Article" {
		t.Fatalf("unexpected title: %q", res.Title)
	}
	if len(res.Images) != 1 {
		t.Fatalf("expected 1 image, got %d", len(res.Images))
	}
	if !strings.Contains(res.Markdown, "images/img_001.") {
		t.Fatalf("markdown image path not rewritten: %q", res.Markdown)
	}
	if memStore.saved == nil {
		t.Fatalf("storage save should be called")
	}
}

type stubDriver struct {
	res *core.FetchResult
	err error
}

func (s *stubDriver) Fetch(_ context.Context, _ string) (*core.FetchResult, error) {
	return s.res, s.err
}
