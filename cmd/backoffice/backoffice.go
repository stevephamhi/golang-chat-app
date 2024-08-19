package backoffice

import (
	"chatgame/config"
	"database/sql"
	"fmt"
)

func Init(db *sql.DB) {
	Run(NewServer(fmt.Sprintf(":%s", config.Env.BackofficeServePort), db))
}
