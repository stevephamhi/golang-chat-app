package backoffice_handlers

import (
	backoffice_controllers "chatgame/modules/backoffice/dashboard/controllers"
	"database/sql"
	"html/template"

	"github.com/gorilla/mux"
)

var (
	dashboardController *backoffice_controllers.DashboardController
)

func init() {
	dashboardController = &backoffice_controllers.DashboardController{
		Templates: template.Must(template.ParseGlob("templates/pages/dashboard/*.html")),
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
	router.HandleFunc("/", dashboardController.Index)
}
