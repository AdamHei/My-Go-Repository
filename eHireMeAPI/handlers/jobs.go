package handlers

import (
	"net/http"
	"github.com/adamhei/eHireMeAPI/models"
	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
	"fmt"
)

// All_jobs will attempt to return all jobs in the database
func (env *Env) All_jobs(writer http.ResponseWriter, req *http.Request) {
	jobs, err := models.All_jobs(env.Db)

	if err != nil {
		respond(writer,
			"We could not fetch all jobs for you",
			&models.My_error{err.Error(), http.StatusInternalServerError})
	} else {
		respond(writer, jobs, nil)
	}
}

// Jobs_by_employer will, given an employer id, return all jobs the employer has created
func (env *Env) Jobs_by_employer(writer http.ResponseWriter, req *http.Request) {
	args := mux.Vars(req)
	emp_id, err := strconv.ParseInt(args["id"], 10, 64)
	if err != nil {
		respond(writer,
			"Give me an employer id!",
			&models.My_error{err.Error(), http.StatusBadRequest})
		return
	}

	jobs, err := models.Get_jobs_by_employer(env.Db, int(emp_id))

	if err != nil {
		respond(writer,
			"We could not fetch all jobs for that employer",
			&models.My_error{err.Error(), http.StatusInternalServerError})
	} else {
		respond(writer, jobs, nil)
	}
}

// Job_id will, given a job id, attempt to return the job
func (env *Env) Job_id(writer http.ResponseWriter, req *http.Request) {
	args := mux.Vars(req)
	job_id, err := strconv.ParseInt(args["id"], 10, 64)
	if err != nil {
		respond(writer, "Give me a job id!", &models.My_error{err.Error(), http.StatusBadRequest})
		return
	}

	jobs, err := models.Get_job(env.Db, int(job_id))

	if err != nil {
		respond(writer,
			"Could not fetch that job",
			&models.My_error{err.Error(), http.StatusInternalServerError})
	} else {
		respond(writer, jobs, nil)
	}
}

// Create_job will, given a job, attempt to store and return the job if successful
func (env *Env) Create_job(writer http.ResponseWriter, req *http.Request) {
	job := new(models.Job)

	err := json.NewDecoder(req.Body).Decode(job)
	if err != nil {
		respond(writer, "Your format sucked", &models.My_error{err.Error(), http.StatusBadRequest})
		return
	}

	job, err = models.Store_job(env.Db, job, false)
	if err != nil {
		respond(writer,
			"Unable to store job",
			&models.My_error{err.Error(), http.StatusInternalServerError})
	} else {
		respond(writer, job, nil)
	}
}

// Update_job will, given a partial job, attempt to update and return the job if successful
func (env *Env) Update_job(writer http.ResponseWriter, req *http.Request) {
	job := new(models.Job)
	err := json.NewDecoder(req.Body).Decode(job)

	if err != nil {
		respond(writer, "Your format sucked", &models.My_error{err.Error(), http.StatusBadRequest})
		return
	}

	job, err = models.Update_job(env.Db, job)
	if err != nil {
		respond(writer,
			"Could not update job",
			&models.My_error{err.Error(), http.StatusInternalServerError})
	} else {
		respond(writer, job, nil)
	}
}

// Delete_job will, given a job id, attempt to delete a job from the database
func (env *Env) Delete_job(writer http.ResponseWriter, req *http.Request) {
	args := mux.Vars(req)
	job_id, err := strconv.ParseInt(args["id"], 10, 64)
	fmt.Printf("Job id: %d was given\n", int(job_id))
	if err != nil {
		respond(writer, "Give me an integer id", &models.My_error{err.Error(), http.StatusBadRequest})
		return
	}

	err = models.Delete_job(env.Db, int(job_id))

	if err != nil {
		respond(writer,
			"We could not delete that job",
			&models.My_error{err.Error(), http.StatusInternalServerError})
	} else {
		respond(writer, "Job deleted", nil)
	}
}
