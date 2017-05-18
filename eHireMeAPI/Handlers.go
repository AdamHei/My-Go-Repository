package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/adamhei/models"
	"log"
	"encoding/json"
	"strconv"
)

func (env *Env) Index(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprint(writer, "Bienvenido!")
}

func (env *Env) AllApplicants(writer http.ResponseWriter, req *http.Request) {
	applicants, err := models.AllApplicants(env.db)

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

	app, err := models.GetApplicant(env.db, int(user_id))

	if err != nil {
		respond(writer, "We were unable to retrieve that applicant", err)
	} else {
		respond(writer, app, nil)
	}
}

func (env *Env) CreateApplicant(w http.ResponseWriter, r *http.Request) {
	app := new(models.Applicant)

	//arr := make([]byte, 1000)
	//r.Body.Read(arr)
	//fmt.Println(string(arr))
	err := json.NewDecoder(r.Body).Decode(app)
	if err != nil {
		respond(w, "Your format sucked", err)
		return
	}

	err = models.StoreApplicant(env.db, app)
	if err != nil {
		respond(w, "Unable to store applicant", err)
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
