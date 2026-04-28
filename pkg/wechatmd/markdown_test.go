package wechatmd

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	prs "github.com/SolaTyolo/wechat-article-markdown/pkg/wechatmd/parser"
)

func TestSelectionToMarkdown(t *testing.T) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(`
<div id="js_content">
  <h2>Title</h2>
  <p>text <strong>bold</strong></p>
  <pre class="language-go"><code>fmt.Println("hi")</code></pre>
  <ul><li>a</li><li>b</li></ul>
</div>`))
	if err != nil {
		t.Fatalf("build doc: %v", err)
	}
	md := prs.SelectionToMarkdown(doc.Find("#js_content"))
	wantTokens := []string{
		"## Title",
		"text **bold**",
		"```go",
		`fmt.Println("hi")`,
		"- a",
		"- b",
	}
	for _, token := range wantTokens {
		if !strings.Contains(md, token) {
			t.Fatalf("markdown missing token %q:\n%s", token, md)
		}
	}
}
