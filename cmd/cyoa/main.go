package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/voroskoi/cyoa/cyoa"
)

func main() {
	story := flag.String("story", "../../gopher.json", "JSON file that contains the story")
	flag.Parse()

	ps := cyoa.ParseJSON(*story)

	cyoahandler := cyoa.NewHandler(ps)

	err := http.ListenAndServe(":8080", cyoahandler)
	if err != nil {
		log.Fatalf("server stopped inexpectedly: %s", err)
	}
}
