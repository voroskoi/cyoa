package cyoa

import (
	"fmt"
	"net/http"
)

type StoryHandler struct {
	story Story
}

func (sh StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL is: %s\n", r.URL.EscapedPath())
	fmt.Fprintln(w, sh.story)
}

func (sh *StoryHandler) SetStory(story Story) {
	sh.story = story
}
