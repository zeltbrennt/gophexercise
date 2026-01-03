package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"adventure/handler"
	"adventure/story"
)

func main() {
	file := flag.String("file", "story.json", "Story in JSON format")
	port := flag.Int("port", 3000, "Port of the server")
	flag.Parse()

	story, err := story.ParseJSON(*file)
	if err != nil {
		log.Fatalf("an error occured: %v", err)
	}

	handler := handler.New(story)
	log.Printf("listening on port :%d...\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handler))
}
