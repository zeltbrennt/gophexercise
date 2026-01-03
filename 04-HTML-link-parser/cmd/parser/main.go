package main

import (
	"flag"
	"fmt"
	"os"

	"linkparser/linkparser"
)

func main() {
	fileName := flag.String("file", "examples/ex1.html", "HTML file to extract links from")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic("file not found")
	}
	links, err := linkparser.Parse(file)
	if err != nil {
		panic(err)
	}
	for _, link := range links {
		fmt.Printf("%+v\n", link)
	}
}
