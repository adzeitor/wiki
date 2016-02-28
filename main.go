package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	storage := NewMemDocuments()

	d1 := storage.Add("1. One")
	d2 := storage.Edit(d1.ID, "2. Two ")
	d3 := storage.Edit(d2.ID, "3. Three")

	pages, _ := NewPageController(storage)
	http.Handle("/page", pages)

	blobs, _ := NewBlobController(storage)
	http.Handle("/blob", blobs)

	fmt.Println("listen http://127.0.0.1:8080")
	fmt.Printf("http://127.0.0.1:8080/page?id=%s\n", d3.ID)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
