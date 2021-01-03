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
	chapter, exist := sh.story[strings.Trim(r.URL.Path, "/")]
	if !exist {
		chapter = sh.story["intro"]
	}
	tmp := template.Must(template.ParseFiles("../../template/template.html"))
	tmp.Execute(w, chapter)
}

func (sh *StoryHandler) SetStory(story Story) {
	sh.story = story
}
