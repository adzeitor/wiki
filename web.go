package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageController struct {
	Wiki     *Wiki
	Template *template.Template
}

func NewPageController(wiki *Wiki) (*PageController, error) {
	tmpl, err := template.ParseFiles("templates/page.html")

	return &PageController{
		Template: tmpl,
		Wiki:     wiki,
	}, err
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
	hash := r.URL.Path

	page := c.Wiki.GetChain(hash)

	c.Template.Execute(w, page)
}

func (c *PageController) Edit(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Path

	r.ParseForm()

	content := r.FormValue("content")

	d := c.Wiki.Fork(hash, content)

	http.Redirect(w, r, "/w/"+d.ID, http.StatusFound)
}
