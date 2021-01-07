package cyoa

import (
	"html"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type storyHandler struct {
	story Story
}

func NewHandler(st Story) http.Handler {
	return storyHandler{st}
}

func (sh storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	chapter, exist := sh.story[strings.Trim(r.URL.Path, "/")]
	if !exist {
		chapter = sh.story["intro"]
	}
	for i, p := range chapter.Story {
		chapter.Story[i] = html.EscapeString(p)
	}
	tmp := template.Must(template.ParseFiles("../../template/template.html"))
	err := tmp.Execute(w, chapter)
	if err != nil {
		log.Printf("unable to execute the template: %s", err)
		http.Error(w, "unable to execute template", http.StatusInternalServerError)
	}
}
