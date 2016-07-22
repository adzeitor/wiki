package main

type MemoryNames struct {
	Items map[string]string
}

func NewMemoryNames() *MemoryNames {
	return &MemoryNames{
		Items: make(map[string]string),
	}
}

func (names *MemoryNames) Link(name string, d Doc) {
	names.Items[name] = d.ID
}

func (names *MemoryNames) GetHash(id string) string {
	return names.Items[id]
}
