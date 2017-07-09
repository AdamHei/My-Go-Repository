package models

import (
	"database/sql"
	"fmt"
	"log"
)

type Job struct {
	ID               int `json:"id"`
	Employer_ID      int `json:"employer_id"`
	Description      string `json:"description"`
	Compensation     string `json:"compensation"`
	Title            string `json:"title"`
	Field            string `json:"field"`
	Title_experience string `json:"title_experience"`
	Field_experience string `json:"field_experience"`
	City             string `json:"city"`
	State            string `json:"state"`
	Active           bool `json:"active"`
}

//Return all jobs in the database, unsorted
func All_jobs(db *sql.DB) ([]*Job, error) {
	return jobs_by_query(db, "SELECT * FROM jobs")
}

//Retrieve all jobs posted by a given employer
func Get_jobs_by_employer(db *sql.DB, employer_id int) ([]*Job, error) {
	query := fmt.Sprintf("SELECT * FROM jobs WHERE employer_id=%d LIMIT 1", employer_id)
	return jobs_by_query(db, query)
}

//Return a list of jobs that satisfy a given query
func jobs_by_query(db *sql.DB, query string) ([]*Job, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	jobs := make([]*Job, 0)
	for rows.Next() {
		job := new(Job)
		err = rows.Scan(job.ID, job.Employer_ID, job.Description, job.Compensation, job.Title, job.Field,
			job.Title_experience, job.Field_experience, job.City, job.State, job.Active)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return jobs, nil
}

//Retrieve a job by its unique id
func Get_job(db *sql.DB, id int) (*Job, error) {
	return job_by_query(db, "SELECT * FROM jobs WHERE id=? LIMIT 1", id)
}

//Return one job that satisfies the given query
func job_by_query(db *sql.DB, query string, param interface{}) (*Job, error) {
	job := new(Job)
	err := db.QueryRow(query, param).Scan(job.ID, job.Employer_ID, job.Description, job.Compensation,
		job.Title, job.Field, job.Title_experience, job.Field_experience, job.City, job.State, job.Active)

	if err != nil {
		return nil, err
	}

	return job, err
}

//TODO Jobs sorted for applicant

//TODO Jobs an applicant has matched with - Match

//TODO Applicants that have applied towards a job - Matches

// Store_job will attempt to insert a given job into the proper table and return it
// withID = false when inserting for the first time
// true when performing a partial update
func Store_job(db *sql.DB, job *Job, withID bool) (*Job, error) {
	query := insert_query(*job, "jobs", withID)
	res, err := db.Exec(query)

	if err != nil {
		return nil, err
	}

	id, _ := res.LastInsertId()
	log.Println("Job stored with id:", id)

	return Get_job(db, int(id))
}

// Update_job will attempt to perform a partial update, by merging the job from the database with the given job
func Update_job(db *sql.DB, job *Job) (*Job, error) {
	if job.ID <= 0 {
		return nil, &my_error{"Send me a valid id"}
	}
	if !job_exists(db, job.ID) {
		return nil, &my_error{"Job doesn't exist"}
	}

	stored_job, err := Get_job(db, job.ID)
	if err != nil {
		return nil, err
	}

	merged_job := merge_jobs(stored_job, job)

	err = Delete_job(db, merged_job.ID)
	if err != nil {
		return nil, err
	}

	return Store_job(db, merged_job, true)
}

//TODO Set active/inactive

// Delete_job will attempt to remove a job by id from the database
func Delete_job(db *sql.DB, id int) error {
	fmt.Println(id)
	res, err := db.Exec("DELETE FROM jobs WHERE id=?", id)
	insert_id, _ := res.LastInsertId()
	log.Printf("Job with id: %d was deleted", insert_id)

	return err
}

//TODO Search
