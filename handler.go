package cyoa

import (
	"fmt"
	"net/http"
	"strings"
)

type StoryHandler struct {
	story Story
}

func (sh StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL is: %s\n", r.URL.EscapedPath())
	if chapter, exist := sh.story[strings.Trim(r.URL.Path, "/")]; exist {
		fmt.Fprintln(w, chapter.Story)
	}
	fmt.Fprintf(w, "\n\n\n%+v", sh.story)
}

func (sh *StoryHandler) SetStory(story Story) {
	sh.story = story
}
