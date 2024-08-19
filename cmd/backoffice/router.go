package backoffice

import (
	"chatgame/config"
	backoffice_handlers "chatgame/modules/backoffice/dashboard/handlers"
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

	subrouter := router.PathPrefix("/backoffice/api/v1").Subrouter()

	subrouter.HandleFunc("/ping", HealthCheckHandler).Methods(http.MethodGet)

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(config.Env.BackoffieStaticDir))))

	backofficePort := ":" + config.Env.BackofficeServePort

	log.Printf("Backoffice server running on %s", backofficePort)
	log.Fatal(http.ListenAndServe(backofficePort, router))

	backoffice_handlers.DashboardHandler(router, s.db)

	return router
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Backoffice PONG...!"))
}
