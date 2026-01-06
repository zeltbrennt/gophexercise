package main

import (
	"flag"
	"log"
	"os"

	sitemapbuilder "sitemap"
)

func main() {
	url := flag.String("url", "http://localhost:3000", "URL of the site to get the sitemap from")

	sb := sitemapbuilder.New(*url)
	err := sb.BuildMap()
	if err != nil {
		log.Fatalf("sitemapbuilder failed: %s", err)
	}
	sb.Write(os.Stdout)
}
