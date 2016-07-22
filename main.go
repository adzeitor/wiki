package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	storage := NewMemoryDocuments()
	names := NewMemoryNames()
	wiki := &Wiki{
		Storage: storage,
		Names: names,
	}
	namesController := NewNamesController(names)

	pages,err := NewPageController(wiki)
	if err != nil {
		log.Fatal(err)
	}

	blobs := NewBlobController(storage)

	createPage, err := NewCreateController(wiki)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/wiki/", http.StripPrefix("/wiki/", namesController))
	http.Handle("/w/", http.StripPrefix("/w/", pages))
	http.Handle("/create/", http.StripPrefix("/create/", createPage))
	http.Handle("/raw/", http.StripPrefix("/raw/", blobs))

	fmt.Println("listen http://127.0.0.1:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
