package database

import (
	"chatgame/config"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMysql(cf mysql.Config) *sql.DB {
	db, err := sql.Open(config.Env.DBDriver, cf.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Connect(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
