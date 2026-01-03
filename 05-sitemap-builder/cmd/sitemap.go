package main

import (
	"fmt"
	"htmlparser/links"
	"strings"
)

func main() {
	links, err := links.GetAll(strings.NewReader(`<a href="/foo">bar</a>`))
	if err != nil {
		panic("!")
	}
	fmt.Println(links)
}
