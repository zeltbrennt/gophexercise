package sitemapbuilder

import (
	"fmt"
	"htmlparser/links"
	"io"
	"log"
	"net/http"
	"strings"

	"sitemap/xml"
)

type SitemapBuilder struct {
	domain string
	links  map[string]bool
	xml    xml.Builder
}

func New(url string) SitemapBuilder {
	links := make(map[string]bool)
	return SitemapBuilder{url, links, xml.Builder{}}
}

func (s *SitemapBuilder) BuildMap() error {
	s.walk(s.domain)
	return nil
}

func (s *SitemapBuilder) walk(next string) {
	s.links[next] = true
	resp, err := http.Get(next)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	links, err := links.GetAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	for _, link := range links {
		var fullLink string
		if strings.HasPrefix(link.Href, s.domain) {
			fullLink = link.Href
		} else if link.Href[0] == '/' {
			fullLink = fmt.Sprintf("%s%s", s.domain, link.Href)
		} else {
			continue
		}
		if !s.links[fullLink] {
			s.walk(fullLink)
		}
	}
}

func (s SitemapBuilder) Write(w io.Writer) {
	s.xml.CreateXML(s.links)
	s.xml.Write(w)
}
