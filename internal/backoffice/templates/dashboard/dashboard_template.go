package dashboard_backoffice_template

import (
	backoffice_dashboard_app "chatgame/internal/backoffice/app/dashboard"
	"html/template"
	"log"
	"net/http"
)

var (
	Template *template.Template
)

func init() {
	var err error
	Template, err = template.ParseGlob("templates/dashboard/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}
}

func RenderIndex(w http.ResponseWriter, vm backoffice_dashboard_app.IndexViewModel) {
	if err := Template.Lookup("index.html").Execute(w, vm); err != nil {
		http.Error(w, "[BACKOFFICE TEMPLATE] Failed to render dashoard index", http.StatusInternalServerError)
	}
}
