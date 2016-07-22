package main

type Wiki struct {
	// Root    Doc
	Storage Storage
	Names   Names
}

type Storage interface {
	Add(Doc) error
	Get(string) (Doc, bool)
}

type hash string

type Names interface {
	GetHash(string) string
	Link(string, Doc)
}

type Page struct {
	Name string
	Doc  Doc
}

func (wiki *Wiki) Create(name string, content string) Doc {
	doc := NewDocument("", "text/plain", content)
	wiki.Storage.Add(doc)
	wiki.Names.Link(name, doc)

	return doc
}

func (wiki *Wiki) Get(h string) Doc {
	doc, _ := wiki.Storage.Get(h)

	return doc
}

func (wiki *Wiki) GetByName(name string) Doc {
	h := wiki.Names.GetHash(name)
	doc, _ := wiki.Storage.Get(string(h))

	return doc
}

func (wiki *Wiki) Fork(parentId string, content string) Doc {
	parent, ok := wiki.Storage.Get(parentId)

	if !ok {
		// FIXME: return error
		return Doc{}
	}
	doc := NewDocument(parent.ID, "text/plain", content)

	wiki.Storage.Add(doc)
	return doc
}

type DocChain struct {
	Doc     Doc
	History []Doc
}

func (wiki *Wiki) GetChain(parentId string) DocChain {
	var items []Doc

	doc, _ := wiki.Storage.Get(parentId)

	id := doc.ID

	for {
		d := wiki.Get(id)
		items = append(items, d)

		if d.Parent == "" {
			break
		}

		id = d.Parent
	}

	chain := DocChain{
		Doc:     doc,
		History: items,
	}

	return chain
}
