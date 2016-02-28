package main

import (
	"crypto/sha512"
	"encoding/hex"
)

type Doc struct {
	Parent      string
	ID          string
	Content     string
	ContentType string
}

type Storage interface {
	Add(string) Doc
	Get(string) Doc
	Edit(string, string) Doc
	GetChain(string) []Doc
}

type hash string

type Names interface {
	Get(string) hash
	Link(hash, string)
}

func NewDocument(parent string, contentType string, content string) Doc {
	h := sha512.New()

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
	}
}
