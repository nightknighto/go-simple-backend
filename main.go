package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var router *mux.Router

func main() {
	db = initializeDatabase(DbUser, DbPassword, DbName)
	router = initializeRouter()

	startServer(":10000", router)
}
