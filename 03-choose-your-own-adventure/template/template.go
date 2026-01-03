package template

import (
	"adventure/story"
	"html/template"
	"os"
)

func Parse(path string) *template.Template {
	raw, err := os.ReadFile(path)
	if err != nil {
		panic("no template file found")
	}

	templ := template.Must(template.New("page").Parse(string(raw)))
	return templ
}

type TemplateData struct {
	story.Chapter
	ID string
}
