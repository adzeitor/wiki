package main

type MemDocuments struct {
	Root  []Doc
	Items map[string]Doc
}

func NewMemDocuments() *MemDocuments {
	return &MemDocuments{
		Root:  nil,
		Items: make(map[string]Doc),
	}
}

func (ds *MemDocuments) Add(content string) Doc {
	doc := NewDocument("", "text/plain", content)
	ds.Root = append(ds.Root, doc)
	ds.Items[doc.ID] = doc

	return doc
}

func (ds *MemDocuments) Edit(id string, content string) Doc {
	parent := ds.Get(id)

	// TODO: content-type
	doc := NewDocument(parent.ID, "text/plain", content)
	ds.Items[doc.ID] = doc

	return doc
}

func (ds *MemDocuments) Get(id string) Doc {
	return ds.Items[id]
}

func (ds *MemDocuments) GetChain(id string) []Doc {
	var items []Doc

	for {
		doc := ds.Items[id]
		items = append(items, doc)

		if doc.Parent == "" {
			return items
		}

		id = doc.Parent
	}

	return items
}
