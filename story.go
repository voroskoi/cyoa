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

var ParsedStory CyoaStory

func ParseJSON(storyfile string) {
	storyfd, err := os.Open(storyfile)
	if err != nil {
		log.Panicf("unable to open story file %q", storyfile)
	}
	decoder := json.NewDecoder(storyfd)
	err = decoder.Decode(&ParsedStory)
	if err != nil {
		log.Panicf("unable to decode: %s", err)
	}
}
