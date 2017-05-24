package handlers

import (
	"net/http"
	"fmt"
	"database/sql"
)

type Env struct {
	Db *sql.DB
}

func (env *Env) Index(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprint(writer, "Bienvenido!")
}
