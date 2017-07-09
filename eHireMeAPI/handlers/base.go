// Package handlers provides controllers for interfacing with all package models methods
package handlers

import (
	"net/http"
	"fmt"
	"database/sql"
	"log"
	"encoding/json"
)

// Type Env serves as a form of dependency injection to neatly deal with passing the database around
// See: http://www.alexedwards.net/blog/organising-database-access
type Env struct {
	Db *sql.DB
}

// Index serves the base route response
func (env *Env) Index(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprint(writer, "Bienvenido!")
}

// respond is a helper function used throughout the package to easily print errors and send responses to the client
func respond(writer http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		log.Println(err.Error())
	}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(data)
}
