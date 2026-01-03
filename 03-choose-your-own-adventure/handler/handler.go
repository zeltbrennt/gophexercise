package handler

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"adventure/story"
	adventureTemplate "adventure/template"
)

var (
	defaultTemplate         *template.Template
	defaultStaticFileServer http.Handler
)

type handler struct {
	Story         story.Story
	Template      *template.Template
	StaticHandler http.Handler
}

func init() {
	defaultTemplate = adventureTemplate.Parse("template/template.html")
	defaultStaticFileServer = http.FileServer(http.Dir("./static"))
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

// Functional Options :)

type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.Template = t
	}
}

func WithStaticFileHandler(sfh http.Handler) HandlerOption {
	return func(h *handler) {
		h.StaticHandler = sfh
	}
}

func New(story story.Story, opts ...HandlerOption) http.Handler {
	handler := handler{Story: story, Template: defaultTemplate, StaticHandler: defaultStaticFileServer}
	for _, opt := range opts {
		opt(&handler)
	}
	return handler
}
