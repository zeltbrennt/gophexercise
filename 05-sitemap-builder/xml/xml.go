package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"slices"
	"strings"
)

const (
	xmlns     = "http://www.sitemaps.org/schemas/sitemap/0.9"
	xmlheader = `<?xml version="1.0" encoding="UTF-8"?>`
)

type url struct {
	Loc string `xml:"loc"`
}

type urlset struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []url    `xml:"url"`
}

type Builder struct {
	output []byte
}

func (b *Builder) CreateXML(links map[string]bool) {
	sitemap := urlset{Xmlns: xmlns}
	urls := make([]url, 0)
	for link := range links {
		urls = append(urls, url{link})
	}
	slices.SortFunc(urls, func(u1, u2 url) int {
		return strings.Compare(u1.Loc, u2.Loc)
	})
	sitemap.URLs = urls
	output, err := xml.MarshalIndent(sitemap, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	b.output = output
}

func (b Builder) Write(writer io.Writer) error {
	var buf bytes.Buffer
	buf.Write([]byte(xmlheader))
	buf.Write([]byte("\n"))
	buf.Write(b.output)
	_, err := buf.WriteTo(writer)
	if err != nil {
		return fmt.Errorf("error writing XML: %v", err)
	}
	return nil
}
