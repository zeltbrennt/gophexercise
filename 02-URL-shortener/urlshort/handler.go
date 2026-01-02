package urlshort

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-yaml/yaml"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url, ok := pathsToUrls[r.URL.Path]; ok {
			log.Println("redirecting to:", url)
			http.Redirect(w, r, url, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathToUrls, err := parseYaml(yml)
	if err != nil {
		return nil, fmt.Errorf("parsing yaml failes: %v", err)
	}
	return MapHandler(pathToUrls, fallback), nil
}

type redirect struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYaml(yml []byte) (map[string]string, error) {
	var redirect []redirect
	err := yaml.Unmarshal(yml, &redirect)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling failed: %v", err)
	}
	asMap := make(map[string]string)
	for _, entry := range redirect {
		asMap[entry.Path] = entry.URL
	}
	return asMap, nil
}
