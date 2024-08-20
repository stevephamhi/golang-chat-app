package backoffice_user_handler

import (
	backoffice_user_app "chatgame/internal/backoffice/app/user"
	backoffice_user_template "chatgame/internal/backoffice/templates/users"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct{}

func NewHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) ViewIndex(w http.ResponseWriter, r *http.Request) {
	vm := backoffice_user_app.IndexViewModel{
		PageTitle: "List Users",
	}

	backoffice_user_template.RenderIndex(w, vm)
}

func (h *UserHandler) ViewCreate(w http.ResponseWriter, r *http.Request) {
	vm := backoffice_user_app.IndexViewModel{
		PageTitle: "Create User",
	}

	backoffice_user_template.RenderCreate(w, vm)
}

func (h *UserHandler) ViewEdit(w http.ResponseWriter, r *http.Request) {
	vm := backoffice_user_app.IndexViewModel{
		PageTitle: "Edit User",
	}

	backoffice_user_template.RenderEdit(w, vm)
}

func (h *UserHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", h.ViewIndex)
	router.HandleFunc("/users/create", h.ViewCreate)
	router.HandleFunc("/users/{id}/edit", h.ViewEdit)
}
