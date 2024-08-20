package backoffice_dashboard_handler

import (
	backoffice_dashboard_app "chatgame/internal/backoffice/app/dashboard"
	dashboard_backoffice_template "chatgame/internal/backoffice/templates/dashboard"
	"net/http"

	"github.com/gorilla/mux"
)

type DashboardHandler struct{}

func NewHandler() *DashboardHandler {
	return &DashboardHandler{}
}

func (h *DashboardHandler) ViewIndex(w http.ResponseWriter, r *http.Request) {
	vm := backoffice_dashboard_app.IndexViewModel{
		PageTitle: "Dashboard",
	}

	dashboard_backoffice_template.RenderIndex(w, vm)
}

func (h *DashboardHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.ViewIndex)
}
