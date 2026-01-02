package main

import (
	"adventure/story"
	adventuretemplate "adventure/template"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Adventurer struct {
	Story         story.Story
	Template      *template.Template
	StaticHandler http.Handler
}

type TemplateData struct {
	story.Chapter
	Id string
}

func (a *Adventurer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		http.Redirect(w, r, "/intro", http.StatusSeeOther)
	}
	if strings.HasPrefix(path, "/static") {
		http.StripPrefix("/static/", a.StaticHandler).ServeHTTP(w, r)
	}
	log.Println(path)
	arc := strings.TrimPrefix(path, "/")
	chapter, ok := a.Story[arc]
	if !ok {
		http.NotFound(w, r)
	}
	err := a.Template.Execute(w, TemplateData{Chapter: chapter, Id: arc})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	story, err := story.ParseJSON("story.json")
	if err != nil {
		log.Fatalf("an error occured: %v", err)
	}
	template, err := adventuretemplate.ParseTemplate("template/template.html")
	if err != nil {
		log.Fatalf("an error occured: %v", err)
	}

	fs := http.FileServer(http.Dir("./template"))
	adventurer := Adventurer{Story: story, Template: template, StaticHandler: fs}
	log.Println("listening on :8090...")
	http.ListenAndServe(":8090", &adventurer)
}
