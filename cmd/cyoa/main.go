package main

import (
	"flag"
	"fmt"

	"github.com/voroskoi/cyoa/cyoa"
)

func main() {
	story := flag.String("story", "../../gopher.json", "JSON file that contains the story")
	flag.Parse()

	cyoa.ParseJSON(*story)
	fmt.Printf("%+v", cyoa.ParsedStory)
}
