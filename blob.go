package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type BlobController struct {
	Storage Storage
}

func NewBlobController(storage Storage) *BlobController {
	return &BlobController{
		Storage: storage,
	}
}

func (c *BlobController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		c.Add(w, r)
	case "GET":
		c.Get(w, r)
	}
}

func (c *BlobController) Add(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	log.Println(r.Header)
	log.Println(r.ContentLength)

	if err != nil {
		log.Println(err)
	}

	id := r.FormValue("id")

	var doc Doc
	if id == "" {
		c.Storage.Add(Doc{
			Content: string(body),
		})
	}

	w.Write([]byte(doc.ID))
}

func (c *BlobController) Get(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	doc, _ := c.Storage.Get(id)

	w.Header().Set("Content-Type", doc.ContentType)
	w.Write([]byte(doc.Content))
}
