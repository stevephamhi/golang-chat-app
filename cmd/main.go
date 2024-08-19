package main

import (
	"chatgame/cmd/backoffice"
	"chatgame/config"
	"chatgame/database"

	"github.com/go-sql-driver/mysql"
)

func main() {
	dbcf := mysql.Config{
		User:                 config.Env.DBUser,
		Passwd:               config.Env.DBPassword,
		Addr:                 config.Env.DBAddress,
		DBName:               config.Env.DBName,
		Net:                  config.Env.DBNet,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db := database.NewMysql(dbcf)
	database.Connect(db)

	backoffice.Init(db)
}
