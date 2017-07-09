package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"github.com/adamhei/eHireMeAPI/handlers"
	"github.com/adamhei/eHireMeAPI/routes"
)

// main, the driver of the eHireMeAPI, opens a database connection, fetches a new Router, and kicks off the HTTP server
func main() {
	dataSourceName := DB_USER + ":" + DB_PASSWORD + "@/eHireMe?parseTime=true"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	env := &handlers.Env{Db: db}

	router := routes.NewRouter(env)
	log.Fatal(http.ListenAndServe(":8080", router))
}
