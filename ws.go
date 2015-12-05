package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	root := "."
	flag.Parse()
	if flag.NArg() > 0 {
		root = flag.Arg(0)
	}
	log.Printf("ws listening on :8180, serving files from %s", root)
	log.Fatal(http.ListenAndServe(":8180", http.FileServer(http.Dir(root))))
}
