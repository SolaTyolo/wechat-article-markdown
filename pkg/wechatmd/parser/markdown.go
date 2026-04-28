package parser

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// SelectionToMarkdown renders a goquery selection into markdown text.
func SelectionToMarkdown(sel *goquery.Selection) string {
	var buf bytes.Buffer
	for _, n := range sel.Nodes {
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			renderMarkdownNode(child, &buf)
		}
	}
	out := strings.TrimSpace(buf.String())
	return normalizeMarkdown(out)
}

func renderMarkdownNode(n *html.Node, buf *bytes.Buffer) {
	if n == nil {
		return
	}
	switch n.Type {
	case html.TextNode:
		text := strings.ReplaceAll(n.Data, "\n", " ")
		text = strings.ReplaceAll(text, "\t", " ")
		if text != "" {
			buf.WriteString(text)
		}
	case html.ElementNode:
		switch n.Data {
		case "h1":
			buf.WriteString("\n# ")
			renderChildren(n, buf)
			buf.WriteString("\n\n")
		case "h2":
			buf.WriteString("\n## ")
			renderChildren(n, buf)
			buf.WriteString("\n\n")
		case "h3":
			buf.WriteString("\n### ")
			renderChildren(n, buf)
			buf.WriteString("\n\n")
		case "h4":
			buf.WriteString("\n#### ")
			renderChildren(n, buf)
			buf.WriteString("\n\n")
		case "h5":
			buf.WriteString("\n##### ")
			renderChildren(n, buf)
			buf.WriteString("\n\n")
		case "h6":
			buf.WriteString("\n###### ")
			renderChildren(n, buf)
			buf.WriteString("\n\n")
		case "p", "section", "article", "div":
			renderChildren(n, buf)
			buf.WriteString("\n\n")
		case "strong", "b":
			buf.WriteString("**")
			renderChildren(n, buf)
			buf.WriteString("**")
		case "em", "i":
			buf.WriteString("*")
			renderChildren(n, buf)
			buf.WriteString("*")
		case "code":
			if n.Parent != nil && n.Parent.Data == "pre" {
				return
			}
			buf.WriteString("`")
			buf.WriteString(strings.TrimSpace(nodeText(n)))
			buf.WriteString("`")
		case "pre":
			lang := codeLanguage(n)
			buf.WriteString("\n```")
			buf.WriteString(lang)
			buf.WriteString("\n")
			buf.WriteString(strings.TrimSpace(nodeText(n)))
			buf.WriteString("\n```\n\n")
		case "blockquote":
			lines := strings.Split(strings.TrimSpace(nodeText(n)), "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line != "" {
					buf.WriteString("> ")
					buf.WriteString(line)
					buf.WriteString("\n")
				}
			}
			buf.WriteString("\n")
		case "a":
			href := attr(n, "href")
			text := strings.TrimSpace(nodeText(n))
			if text == "" {
				text = href
			}
			if href == "" {
				buf.WriteString(text)
			} else {
				buf.WriteString("[")
				buf.WriteString(text)
				buf.WriteString("](")
				buf.WriteString(href)
				buf.WriteString(")")
			}
		case "img":
			src := attr(n, "data-src")
			if src == "" {
				src = attr(n, "src")
			}
			alt := attr(n, "alt")
			buf.WriteString("![")
			buf.WriteString(alt)
			buf.WriteString("](")
			buf.WriteString(src)
			buf.WriteString(")\n\n")
		case "ul":
			for child := n.FirstChild; child != nil; child = child.NextSibling {
				if child.Type == html.ElementNode && child.Data == "li" {
					buf.WriteString("- ")
					buf.WriteString(strings.TrimSpace(nodeText(child)))
					buf.WriteString("\n")
				}
			}
			buf.WriteString("\n")
		case "ol":
			index := 1
			for child := n.FirstChild; child != nil; child = child.NextSibling {
				if child.Type == html.ElementNode && child.Data == "li" {
					buf.WriteString(strconv.Itoa(index))
					buf.WriteString(". ")
					buf.WriteString(strings.TrimSpace(nodeText(child)))
					buf.WriteString("\n")
					index++
				}
			}
			buf.WriteString("\n")
		case "br":
			buf.WriteString("\n")
		default:
			renderChildren(n, buf)
		}
	}
}

func renderChildren(n *html.Node, buf *bytes.Buffer) {
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		renderMarkdownNode(child, buf)
	}
}

func nodeText(n *html.Node) string {
	if n == nil {
		return ""
	}
	var buf bytes.Buffer
	var walk func(*html.Node)
	walk = func(cur *html.Node) {
		if cur.Type == html.TextNode {
			buf.WriteString(cur.Data)
		}
		for c := cur.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	return strings.TrimSpace(buf.String())
}

func attr(n *html.Node, key string) string {
	for _, a := range n.Attr {
		if a.Key == key {
			return strings.TrimSpace(a.Val)
		}
	}
	return ""
}

func normalizeMarkdown(s string) string {
	lines := strings.Split(s, "\n")
	out := make([]string, 0, len(lines))
	blank := false
	for _, line := range lines {
		t := strings.TrimRight(line, " \t")
		if strings.TrimSpace(t) == "" {
			if !blank {
				out = append(out, "")
			}
			blank = true
			continue
		}
		blank = false
		out = append(out, t)
	}
	return strings.TrimSpace(strings.Join(out, "\n")) + "\n"
}

func codeLanguage(pre *html.Node) string {
	className := attr(pre, "class")
	if strings.Contains(className, "language-") {
		parts := strings.Split(className, "language-")
		if len(parts) > 1 {
			lang := strings.Fields(parts[1])
			if len(lang) > 0 {
				return lang[0]
			}
		}
	}
	return ""
}
