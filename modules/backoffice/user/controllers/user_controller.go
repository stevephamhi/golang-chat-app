package backoffice_controllers

import (
	"html/template"
	"net/http"
)

type UserController struct {
	Templates *template.Template
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "User Index",
	}

	err := c.Templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "User Create",
	}

	err := c.Templates.ExecuteTemplate(w, "create.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *UserController) Edit(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "User Edit",
	}

	err := c.Templates.ExecuteTemplate(w, "edit.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
