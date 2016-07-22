package main

import (
	"testing"
)

func TestParentMustExist(t *testing.T) {
	wiki := &Wiki{
		Storage: NewMemoryDocuments(),
		Names:   NewMemoryNames(),
	}

	a1 := wiki.Edit("NON_EXISTING_HASH010203", "content")
	a2 := wiki.Get(a1.ID)

	if len(a2.Content) > 0 {
		t.Error("Error parent hash. Document cannot have unexisted parent hash.")
	}
}

func TestForkChain(t *testing.T) {
	contents := []string{"zero", "one", "two", "three"}

	wiki := &Wiki{
		Storage: NewMemoryDocuments(),
		Names:   NewMemoryNames(),
	}
	orig := wiki.Create("doc", contents[0])

	last := orig
	for _, newContent := range contents[1:] {
		last = wiki.Fork(last.ID, newContent)
	}

	for i, x := range wiki.GetChain(last.ID).History {
		c1 := contents[len(contents)-i-1]
		c2 := x.Content

		if c1 != c2 {
			t.Error("Expected ", c1, ", got ", c2)
		}
	}
}
