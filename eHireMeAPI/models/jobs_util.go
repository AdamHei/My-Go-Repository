package models

import (
	"database/sql"
	"fmt"
)

func (job Job) member_fields(withId bool) string {
	fields := "employer_id, description, compensation, title, field, title_experience, field_experience, city, state, active)"
	if withId {
		fields = "(id, " + fields
	} else {
		fields = "(" + fields
	}
	return fields
}

func (job Job) member_values(withId bool) string {
	values := fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)",
		job.Employer_ID, job.Description, job.Compensation, job.Title, job.Field,
		job.Title_experience, job.Field_experience, job.City, job.State, job.Active)
	if withId {
		return add_ID(values, job.ID)
	} else {
		return "(" + values
	}
}

func job_exists(db *sql.DB, query_id int) bool {
	id := -1
	err := db.QueryRow("SELECT id FROM jobs WHERE id=? LIMIT 1", query_id).Scan(&id)

	return id != -1 || err == nil
}

func merge_jobs(originalJob, newJob *Job) *Job {
	merged := new(Job)
	merged.ID = originalJob.ID
	merged.Employer_ID = choose_int(originalJob.Employer_ID, newJob.Employer_ID)
	merged.Description = choose_string(originalJob.Description, newJob.Description)
	merged.Compensation = choose_string(originalJob.Compensation, newJob.Compensation)
	merged.Title = choose_string(originalJob.Title, newJob.Title)
	merged.Field = choose_string(originalJob.Field, newJob.Field)
	merged.Title_experience = choose_string(originalJob.Title_experience, newJob.Title_experience)
	merged.Field_experience = choose_string(originalJob.Field_experience, newJob.Field_experience)
	merged.City = choose_string(originalJob.City, newJob.City)
	merged.State = choose_string(originalJob.State, newJob.State)
	merged.Active = originalJob.Active || newJob.Active

	return merged
}
