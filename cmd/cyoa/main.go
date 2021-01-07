package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/voroskoi/cyoa"
)

var (
	story = flag.String("story", "../../gopher.json", "JSON file that contains the story")
	web   = flag.Bool("web", false, "Run web werson, not the CMD based.")
)

func main() {
	flag.Parse()

	ps := cyoa.ParseJSON(*story)

	if *web {
		cyoahandler := cyoa.NewHandler(ps)

		err := http.ListenAndServe(":8080", cyoahandler)
		if err != nil {
			log.Fatalf("server stopped inexpectedly: %s", err)
		}
	} else {
		printStory(ps, "intro")
	}
}

func printStory(ps cyoa.Story, chapter string) {
	ch := ps[chapter]
	fmt.Printf("Title: %s\n\n", ch.Title)
	for _, p := range ch.Story {
		fmt.Println(p)
	}
	for i := range ch.Options {
		fmt.Printf("\nPress %d to continue with: %s\n", i, ch.Options[i].Text)
	}
	if len(ch.Options) == 0 {
		return
	}
	var jump int
	fmt.Scanln(&jump)
	printStory(ps, ch.Options[jump].Arc)
}
