package backoffice_handlers

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func DashboardHandler(router *mux.Router, db *sql.DB) {

}
