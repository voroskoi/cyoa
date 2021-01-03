package cyoa

import (
	"encoding/json"
	"log"
	"os"
)

type CyoaStory map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func ParseJSON(storyfile string) CyoaStory {
	storyfd, err := os.Open(storyfile)
	if err != nil {
		log.Panicf("unable to open story file %q", storyfile)
	}
	decoder := json.NewDecoder(storyfd)
	var story CyoaStory
	err = decoder.Decode(&story)
	if err != nil {
		log.Panicf("unable to decode: %s", err)
	}
	return story
}
