package handlers

import (
	"net/http"
	"fmt"
	"database/sql"
	"log"
	"encoding/json"
)

type Env struct {
	Db *sql.DB
}

func (env *Env) Index(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprint(writer, "Bienvenido!")
}

func respond(writer http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		log.Println(err.Error())
	}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(data)
}
