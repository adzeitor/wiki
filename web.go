package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageController struct {
	Storage  Storage
	Template *template.Template
}

func NewPageController(storage Storage) (*PageController, error) {
	tmpl, err := template.ParseFiles("templates/page.html")

	return &PageController{
		Template: tmpl,
		Storage:  storage,
	}, err
}

type WebWikiPage struct {
	Doc     Doc
	History []Doc
}

func (c *PageController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		c.Edit(w, r)
	case "GET":
		c.Show(w, r)
	default:
		log.Println("other", r.Method)
	}
}

func (c *PageController) Show(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	page := WebWikiPage{
		Doc:     c.Storage.Get(id),
		History: c.Storage.GetChain(id),
	}

	c.Template.Execute(w, page)
}

func (c *PageController) Edit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id := r.FormValue("id")
	content := r.FormValue("content")

	d := c.Storage.Edit(id, content)

	http.Redirect(w, r, "/page?id="+d.ID, http.StatusFound)
}
