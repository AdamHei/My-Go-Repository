package main

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type Env struct {
	db *sql.DB
}

func main() {
	db, err := sql.Open("mysql", "root:Spyrohurricane17@/eHireMe?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	env := &Env{db}

	router := NewRouter(env)
	log.Fatal(http.ListenAndServe(":8080", router))
}
