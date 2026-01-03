package handler

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"adventure/story"
	adventureTemplate "adventure/template"
)

type handler struct {
	Story         story.Story
	Template      *template.Template
	StaticHandler http.Handler
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" || path == "" {
		http.Redirect(w, r, "/intro", http.StatusSeeOther)
		return
	}
	if strings.HasPrefix(path, "/static/") {
		http.StripPrefix("/static/", h.StaticHandler).ServeHTTP(w, r)
		return
	}
	log.Println(path)
	chapter := strings.TrimPrefix(path, "/")
	data, ok := h.Story[chapter]
	if !ok {
		http.NotFound(w, r)
		return
	}
	err := h.Template.Execute(w, adventureTemplate.TemplateData{Chapter: data, ID: chapter})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func New(story story.Story, staticFileHandler http.Handler) http.Handler {
	template := adventureTemplate.Parse("template/template.html")
	return handler{Story: story, Template: template, StaticHandler: staticFileHandler}
}
