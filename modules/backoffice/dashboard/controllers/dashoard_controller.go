package backoffice_controllers

import (
	"html/template"
	"net/http"
)

type DashboardController struct {
	Templates *template.Template
}

func (c *DashboardController) Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Dashboard Index",
	}

	err := c.Templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
