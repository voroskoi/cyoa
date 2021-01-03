package cyoa

import (
	"net/http"
	"strings"
	"text/template"
)

type StoryHandler struct {
	story Story
}

func (sh StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if chapter, exist := sh.story[strings.Trim(r.URL.Path, "/")]; exist {
		tmp := template.Must(template.ParseFiles("../../template/template.html"))
		tmp.Execute(w, chapter)
	}
}

func (sh *StoryHandler) SetStory(story Story) {
	sh.story = story
}
