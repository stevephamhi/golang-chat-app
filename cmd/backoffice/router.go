package backoffice

import (
	"chatgame/config"
	backoffice_dashboard_handler "chatgame/internal/backoffice/handlers/dashboard"
	backoffice_user_handler "chatgame/internal/backoffice/handlers/user"
	"database/sql"
	"html/template"
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

var errorTemplates *template.Template

func init() {
	var err error
	errorTemplates, err = template.ParseGlob("templates/errors/*.html")
	if err != nil {
		log.Fatalf("Failed to parse error templates: %v", err)
	}
}

func Run(s *BackofficeServer) *mux.Router {
	router := mux.NewRouter()

	websubrouter := router.PathPrefix("/").Subrouter()

	backoffice_dashboard_handler.NewHandler().RegisterRoutes(websubrouter)
	backoffice_user_handler.NewHandler().RegisterRoutes(websubrouter)

	// Set 404 template
	websubrouter.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(config.Env.BackoffieStaticDir))))

	log.Printf("Backoffice server running on %s", config.Env.BackofficeServePort)
	log.Fatal(http.ListenAndServe(config.Env.BackofficeServePort, router))

	return router
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	if err := errorTemplates.Lookup("404.html").Execute(w, nil); err != nil {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
	}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Backoffice PONG...!"))
}
