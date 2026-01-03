package linkparser

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Link struct {
	Href string
	Text string
}

func Parse(file io.Reader) ([]Link, error) {
	doc, err := html.Parse(file)
	if err != nil {
		return nil, fmt.Errorf("error parsing html: %s", err)
	}
	links := make([]Link, 0)
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			newLink := Link{}
			newLink.Href = getHrefValue(n)
			newLink.Text = getInnerHTMLText(n)
			links = append(links, newLink)
		}
	}

	return links, nil
}

func getHrefValue(n *html.Node) string {
	for _, a := range n.Attr {
		if a.Key == "href" {
			return a.Val
		}
	}
	return ""
}

func getInnerHTMLText(n *html.Node) string {
	var builder strings.Builder

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.TextNode {
			builder.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	return strings.Join(strings.Fields(builder.String()), " ")
}
