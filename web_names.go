package main

import "net/http"

type NamesController struct {
	Names Names
}

func NewNamesController(names Names) *NamesController {
	return &NamesController{
		Names: names,
	}
}

func (c *NamesController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path
	hash := c.Names.GetHash(name)

	if hash == "" {
		http.Redirect(w, r, "/create/"+name, http.StatusFound)
		return
	}

	http.Redirect(w, r, "/w/"+name+"/"+string(hash), http.StatusFound)
}
