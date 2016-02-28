package main

import (
	"net/http"
)

type BlobController struct {
	Storage Storage
}

func NewBlobController(storage Storage) (*BlobController, error) {
	return &BlobController{
		Storage: storage,
	}, nil
}

func (c *BlobController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	doc := c.Storage.Get(id)

	w.Header().Set("Content-Type", doc.ContentType)
	w.Write([]byte(doc.Content))
}
