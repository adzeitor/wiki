package main

import (
	"testing"
)

func TestChain(t *testing.T) {
	contents := []string{"zero", "one", "two", "three"}

	documents := NewMemDocuments()
	orig := documents.Add(contents[0])

	last := orig
	for _, newContent := range contents[1:] {
		last = documents.Modify(last.ID, newContent)
	}

	for i, x := range documents.GetChain(last.ID) {
		c1 := contents[len(contents)-i-1]
		c2 := x.Content

		if c1 != c2 {
			t.Error("Expected ", c1, ", got ", c2)
		}
	}
}
