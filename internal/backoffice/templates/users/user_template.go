package backoffice_user_template

import (
	backoffice_user_app "chatgame/internal/backoffice/app/user"
	"html/template"
	"log"
	"net/http"
)

var (
	Template *template.Template
)

func init() {
	var err error
	Template, err = template.ParseGlob("templates/users/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}
}

func RenderIndex(w http.ResponseWriter, vm backoffice_user_app.IndexViewModel) {
	if err := Template.Lookup("index.html").Execute(w, vm); err != nil {
		http.Error(w, "[BACKOFFICE TEMPLATE] Failed to render user index", http.StatusInternalServerError)
	}
}

func RenderCreate(w http.ResponseWriter, vm backoffice_user_app.IndexViewModel) {
	if err := Template.Lookup("create.html").Execute(w, vm); err != nil {
		http.Error(w, "[BACKOFFICE TEMPLATE] Failed to render user create", http.StatusInternalServerError)
	}
}

func RenderEdit(w http.ResponseWriter, vm backoffice_user_app.IndexViewModel) {
	if err := Template.Lookup("edit.html").Execute(w, vm); err != nil {
		http.Error(w, "[BACKOFFICE TEMPLATE] Failed to render user edit", http.StatusInternalServerError)
	}
}
