package links

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

func GetAll(file io.Reader) ([]Link, error) {
	doc, err := html.Parse(file)
	if err != nil {
		return nil, fmt.Errorf("error parsing html: %s", err)
	}
	links := make([]Link, 0)
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			newLink := Link{}
			newLink.Href = getHrefValue(n)
			newLink.Text = getInnerText(n)
			links = append(links, newLink)
		}
	}
	return links, nil
}

func getHrefValue(node *html.Node) string {
	for _, a := range node.Attr {
		if a.Key == "href" {
			return a.Val
		}
	}
	return ""
}

func getInnerText(node *html.Node) string {
	var builder strings.Builder
	for n := range node.Descendants() {
		if n.Type == html.TextNode {
			builder.WriteString(n.Data)
		}
	}
	return strings.Join(strings.Fields(builder.String()), " ")
}
