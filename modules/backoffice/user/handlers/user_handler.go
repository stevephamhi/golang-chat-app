package backoffice_handlers

import (
	backoffice_controllers "chatgame/modules/backoffice/user/controllers"
	"database/sql"
	"html/template"

	"github.com/gorilla/mux"
)

var (
	userController *backoffice_controllers.UserController
)

func init() {
	userController = &backoffice_controllers.UserController{
		Templates: template.Must(template.ParseGlob("templates/pages/users/*.html")),
	}
}

type Store struct {
	db *sql.DB
}

type Handler struct {
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/user", userController.Index)
	router.HandleFunc("/user/create", userController.Create)
	router.HandleFunc("/user/edit", userController.Edit)
}
