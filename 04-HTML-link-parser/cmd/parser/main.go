package main

import (
	"flag"
	"fmt"
	"os"

	"htmlparser/links"
)

func main() {
	fileName := flag.String("file", "examples/ex1.html", "HTML file to extract links from")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic("file not found")
	}
	links, err := links.GetAll(file)
	if err != nil {
		panic(err)
	}
	for _, link := range links {
		fmt.Printf("%+v\n", link)
	}
}
