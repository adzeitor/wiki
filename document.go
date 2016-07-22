package main

import (
	"crypto/sha1"
	"encoding/hex"
)

type Doc struct {
	Parent      string
	ID          string
	Content     string
	ContentType string
	Size        int
}

type Document struct {
}

func NewDocument(parent string, contentType string, content string) Doc {
	h := sha1.New()

	// FIXME: is proof of work needed here?
	h.Write([]byte(parent))
	h.Write([]byte(contentType))
	h.Write([]byte(content))

	hash := hex.EncodeToString(h.Sum(nil))

	return Doc{
		Parent:      parent,
		ID:          hash,
		Content:     content,
		ContentType: contentType,
		Size:        len(content),
	}
}
