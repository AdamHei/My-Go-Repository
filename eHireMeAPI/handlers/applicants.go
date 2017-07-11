package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/adamhei/eHireMeAPI/models"
	"encoding/json"
	"strconv"
)

// All_applicants will attempt to return every applicant in the database
func (env *Env) All_applicants(writer http.ResponseWriter, req *http.Request) {
	applicants, err := models.All_applicants(env.Db)

	if err != nil {
		respond(writer,
			"We could not fetch all applicants for you :(",
			&models.My_error{err.Error(), http.StatusInternalServerError}	)
	} else {
		respond(writer, applicants, nil)
	}
}

// Applicant_id will attempt to return an applicant by their id in the database
func (env *Env) Applicant_id(writer http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	user_id, err := strconv.ParseInt(args["id"], 10, 64)
	if err != nil {
		respond(writer,
			"Give me an integer id!",
			&models.My_error{err.Error(), http.StatusBadRequest})
		return
	}

	app, err := models.Get_applicant(env.Db, int(user_id))

	if err != nil {
		respond(writer,
			"We were unable to retrieve that applicant",
			&models.My_error{err.Error(), http.StatusInternalServerError})
	} else {
		respond(writer, app, nil)
	}
}

// Create_applicant will, given an applicant in JSON form, attempt to store the applicant in the database
// and return it if successful
func (env *Env) Create_applicant(w http.ResponseWriter, r *http.Request) {
	app := new(models.Applicant)

	err := json.NewDecoder(r.Body).Decode(app)
	if err != nil {
		respond(w, "Your format sucked", &models.My_error{err.Error(), http.StatusBadRequest})
		return
	}

	app, err = models.Store_applicant(env.Db, app)
	if err != nil {
		respond(w,
			"Unable to store applicant",
			&models.My_error{err.Error(), http.StatusInternalServerError})
	} else {
		respond(w, app, nil)
	}
}

// Update_applicant will, given a partial applicant, attempt to store it in the database and return it if successful
func (env *Env) Update_applicant(w http.ResponseWriter, r *http.Request) {
	app := new(models.Applicant)
	err := json.NewDecoder(r.Body).Decode(app)
	if err != nil {
		respond(w, "Your format sucked", &models.My_error{err.Error(), http.StatusBadRequest})
		return
	}

	app, err = models.Update_applicant(env.Db, app)
	if err != nil {
		respond(w,
			"Could not update applicant",
			&models.My_error{err.Error(), http.StatusInternalServerError})
	} else {
		respond(w, app, nil)
	}
}
