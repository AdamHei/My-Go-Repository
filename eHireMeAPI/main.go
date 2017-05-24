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
	db, err := sql.Open("mysql", "root:Spyrohurricane17@/eHireMe?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	env := &handlers.Env{Db: db}

	router := routes.NewRouter(env)
	log.Fatal(http.ListenAndServe(":8080", router))
}
