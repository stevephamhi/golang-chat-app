package backoffice

import (
	"chatgame/config"
	backoffice_dashboard_handler "chatgame/internal/backoffice/handlers/dashboard"
	backoffice_user_handler "chatgame/internal/backoffice/handlers/user"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type BackofficeServer struct {
	addr string
	db   *sql.DB
}

func NewServer(addr string, db *sql.DB) *BackofficeServer {
	return &BackofficeServer{
		addr: addr,
		db:   db,
	}
}

func Run(s *BackofficeServer) *mux.Router {
	router := mux.NewRouter()

	websubrouter := router.PathPrefix("/").Subrouter()

	backoffice_dashboard_handler.NewHandler().RegisterRoutes(websubrouter)
	backoffice_user_handler.NewHandler().RegisterRoutes(websubrouter)

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(config.Env.BackoffieStaticDir))))

	backofficePort := ":" + config.Env.BackofficeServePort

	log.Printf("Backoffice server running on %s", backofficePort)
	log.Fatal(http.ListenAndServe(backofficePort, router))

	return router
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Backoffice PONG...!"))
}
