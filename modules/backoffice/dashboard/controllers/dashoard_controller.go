package backoffice_controllers

import (
	"net/http"
	"text/template"
)

type DashboardController struct {
	Templates *template.Template
}

func (c *DashboardController) Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"title": "Dashboard",
	}

	err := c.Templates.ExecuteTemplate(w, "backoffice/pages/dashboard/index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
