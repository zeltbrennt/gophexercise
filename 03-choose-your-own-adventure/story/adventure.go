package story

import (
	"encoding/json"
	"fmt"
	"os"
)

type Story map[string]Chapter

type Chapter struct {
	Title   string    `json:"title"`
	Text    []string  `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Link string `json:"arc"`
}

func ParseJSON(path string) (Story, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return Story{}, fmt.Errorf("error reading file: %v", err)
	}
	var story Story
	err = json.Unmarshal(file, &story)
	if err != nil {
		return Story{}, fmt.Errorf("error unmarshalling json: %v", err)
	}
	return story, nil
}
