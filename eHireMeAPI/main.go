package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"github.com/adamhei/handlers"
	"github.com/adamhei/routes"
)

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
