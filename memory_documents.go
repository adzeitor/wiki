package main

type MemoryDocuments struct {
	Root  []Doc
	Items map[string]Doc
}

func NewMemoryDocuments() *MemoryDocuments {
	return &MemoryDocuments{
		Root:  nil,
		Items: make(map[string]Doc),
	}
}

func (ds *MemoryDocuments) Add(doc Doc) error {
	ds.Root = append(ds.Root, doc)
	ds.Items[doc.ID] = doc

	return nil
}

func (ds *MemoryDocuments) Get(id string) (Doc, bool) {
	r, ok := ds.Items[id]

	return r, ok
}
