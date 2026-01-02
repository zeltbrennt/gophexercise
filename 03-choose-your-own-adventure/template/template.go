package template

import (
	"fmt"
	"html/template"
	"os"
)

func ParseTemplate(path string) (*template.Template, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading template file: %v", err)
	}

	templ, err := template.New("page").Parse(string(raw))
	if err != nil {
		return nil, fmt.Errorf("error creating template: %v", err)
	}
	return templ, nil
}
