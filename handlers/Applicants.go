package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/adamhei/models"
	"log"
	"encoding/json"
	"strconv"
)

func (env *Env) AllApplicants(writer http.ResponseWriter, req *http.Request) {
	applicants, err := models.AllApplicants(env.Db)

	if err != nil {
		respond(writer, "We could not fetch all applicants for you :(", err)
	} else {
		respond(writer, applicants, nil)
	}
}

func (env *Env) ApplicantId(writer http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	user_id, err := strconv.ParseInt(args["id"], 10, 64)
	if err != nil {
		respond(writer, "Give me an integer id!", err)
		return
	}

	app, err := models.GetApplicant(env.Db, int(user_id))

	if err != nil {
		respond(writer, "We were unable to retrieve that applicant", err)
	} else {
		respond(writer, app, nil)
	}
}

func (env *Env) CreateApplicant(w http.ResponseWriter, r *http.Request) {
	app := new(models.Applicant)

	err := json.NewDecoder(r.Body).Decode(app)
	if err != nil {
		respond(w, "Your format sucked", err)
		return
	}

	err = models.StoreApplicant(env.Db, app)
	if err != nil {
		respond(w, "Unable to store applicant", err)
	} else {
		respond(w, "Success", nil)
	}
}

func (env *Env) UpdateApplicant(w http.ResponseWriter, r *http.Request) {
	app := new(models.Applicant)
	err := json.NewDecoder(r.Body).Decode(app)
	if err != nil {
		respond(w, "Your format sucked", err)
		return
	}

	err = models.UpdateApplicant(env.Db, app)
	if err != nil {
		respond(w, "Could not update applicant", err)
	} else {
		respond(w, "Success", nil)
	}
}

func respond(writer http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		log.Println(err.Error())
	}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(data)
}
