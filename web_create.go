package main

import (
	"html/template"
	"log"
	"net/http"
)

type CreateController struct {
	Wiki     *Wiki
	Template *template.Template
}

func NewCreateController(wiki *Wiki) (*CreateController, error) {
	tmpl, err := template.ParseFiles("templates/create.html")

	return &CreateController{
		Template: tmpl,
		Wiki:     wiki,
	}, err
}

func (c *CreateController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		c.Create(w, r)
	case "GET":
		c.Form(w, r)
	}

}

func (c *CreateController) Form(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path
	hash := c.Wiki.Names.GetHash(name)

	if hash != "" {
		http.Redirect(w, r, "/wiki/"+name, http.StatusFound)
		return
	}

	c.Template.Execute(w, name)
}

func (c *CreateController) Create(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path

	log.Println(r.URL.Path)

	r.ParseForm()
	content := r.FormValue("content")
	d := c.Wiki.Create(name, content)

	http.Redirect(w, r, "/w/"+name+"/"+d.ID, http.StatusFound)
}
