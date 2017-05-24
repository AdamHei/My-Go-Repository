package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/adamhei/models"
	"encoding/json"
	"strconv"
)

func (env *Env) All_applicants(writer http.ResponseWriter, req *http.Request) {
	applicants, err := models.All_applicants(env.Db)

	if err != nil {
		respond(writer, "We could not fetch all applicants for you :(", err)
	} else {
		respond(writer, applicants, nil)
	}
}

func (env *Env) Applicant_id(writer http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	user_id, err := strconv.ParseInt(args["id"], 10, 64)
	if err != nil {
		respond(writer, "Give me an integer id!", err)
		return
	}

	app, err := models.Get_applicant(env.Db, int(user_id))

	if err != nil {
		respond(writer, "We were unable to retrieve that applicant", err)
	} else {
		respond(writer, app, nil)
	}
}

func (env *Env) Create_applicant(w http.ResponseWriter, r *http.Request) {
	app := new(models.Applicant)

	err := json.NewDecoder(r.Body).Decode(app)
	if err != nil {
		respond(w, "Your format sucked", err)
		return
	}

	err = models.Store_applicant(env.Db, app)
	if err != nil {
		respond(w, "Unable to store applicant", err)
	} else {
		respond(w, "Success", nil)
	}
}

func (env *Env) Update_applicant(w http.ResponseWriter, r *http.Request) {
	app := new(models.Applicant)
	err := json.NewDecoder(r.Body).Decode(app)
	if err != nil {
		respond(w, "Your format sucked", err)
		return
	}

	err = models.Update_applicant(env.Db, app)
	if err != nil {
		respond(w, "Could not update applicant", err)
	} else {
		respond(w, "Success", nil)
	}
}
