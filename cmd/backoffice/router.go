package backoffice

import (
	"chatgame/config"
	dashboard_handlers "chatgame/modules/backoffice/dashboard/handlers"
	user_handlers "chatgame/modules/backoffice/user/handlers"
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

	dashboardHandler := dashboard_handlers.NewHandler(s.db)
	dashboardHandler.RegisterRoutes(websubrouter)

	userHandler := user_handlers.NewHandler(s.db)
	userHandler.RegisterRoutes(websubrouter)

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
