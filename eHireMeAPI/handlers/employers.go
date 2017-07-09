package handlers

import (
	"net/http"
	"github.com/adamhei/eHireMeAPI/models"
	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
)

// All_employers will attempt to return all employers in the database
func (env *Env) All_employers(writer http.ResponseWriter, req *http.Request) {
	employers, err := models.All_employers(env.Db)

	if err != nil {
		respond(writer, "We could not fetch all employers", err)
	} else {
		respond(writer, employers, nil)
	}
}

// Employer_id will attempt to return a given employer by its id in the database
func (env *Env) Employer_id(writer http.ResponseWriter, req *http.Request) {
	args := mux.Vars(req)
	emp_id, err := strconv.ParseInt(args["id"], 10, 64)
	if err != nil {
		respond(writer, "Give me an integer id", err)
		return
	}

	emp, err := models.Get_employer(env.Db, int(emp_id))

	if err != nil {
		respond(writer, "We were unable to retrieve that employer", err)
	} else {
		respond(writer, emp, nil)
	}
}

// Create_employer will, given an employer, attempt to store it in the database and return it if successful
func (env *Env) Create_employer(w http.ResponseWriter, r *http.Request) {
	emp := new(models.Employer)

	err := json.NewDecoder(r.Body).Decode(emp)
	if err != nil {
		respond(w, "Your format sucked", err)
	}

	emp, err = models.Store_employer(env.Db, emp, false)
	if err != nil {
		respond(w, "Unable to store employer", err)
	} else {
		respond(w, emp, nil)
	}
}

// Update_employer will, given a partial employer, attempt to update that employer in the database and return it
// if successful
func (env *Env) Update_employer(w http.ResponseWriter, r *http.Request) {
	emp := new(models.Employer)
	err := json.NewDecoder(r.Body).Decode(emp)
	if err != nil {
		respond(w, "Your format sucked", err)
		return
	}

	emp, err = models.Update_employer(env.Db, emp)
	if err != nil {
		respond(w, "Could not update employer", err)
	} else {
		respond(w, emp, nil)
	}
}
