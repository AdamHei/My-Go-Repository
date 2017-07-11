package models

import (
	"database/sql"
	"fmt"
)

const app_insert_statement = "INSERT INTO applicants " +
	"(name, email, password, dob, age, bio, city, state, title, field, title_experience, field_experience, prof_pic) " +
	"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
const app_update_statement = "UPDATE applicants SET name = ?, email = ?, password = ?, dob = ?, age = ?, bio = ?, " +
	"city = ?, state = ?, title = ?, field = ?, title_experience = ?, field_experience = ?, prof_pic = ? WHERE id = ?"

// model_values, part of the Model interface implementation, returns the values of the receiver Applicant
// corresponding to the Applicant schema, with or without the primary key id
func (app Applicant) model_values(withID bool) []interface{} {
	values := []interface{}{app.Name, app.Email, app.Password, app.Dob.Format("2006-01-02"), app.Age, app.Bio,
							app.City, app.State, app.Title, app.Field, app.Title_Experience,
							app.Field_Experience, app.Prof_Pic_Url}
	if withID {
		return append(values, app.ID)
	} else {
		return values
	}
}

// applicant_exists will, given an applicant id, return whether the applicant is in the database
func applicant_exists(db *sql.DB, query_id int) bool {
	id := -1
	err := db.QueryRow("SELECT id FROM applicants WHERE id=?", query_id).Scan(&id)

	return id != -1 || err == nil
}

// merge_applicants will, given an applicant from the database and one from the client, return a merged version
// for use in a partial update
// Note: IDs are primary keys and cannot be updated and passwords will be updated in a different function
func merge_applicants(origApp, newApp *Applicant) *Applicant {
	merged := new(Applicant)
	//TODO Change password in a diff function
	//ID is immutable; Create new user
	merged.ID = origApp.ID
	merged.Password = origApp.Password
	merged.Name = choose_string(origApp.Name, newApp.Name)
	merged.Email = choose_string(origApp.Email, newApp.Email)
	merged.Bio = choose_string(origApp.Bio, newApp.Bio)
	merged.City = choose_string(origApp.City, newApp.City)
	merged.State = choose_string(origApp.State, newApp.State)
	merged.Title = choose_string(origApp.Title, newApp.Title)
	merged.Field = choose_string(origApp.Field, newApp.Field)
	merged.Title_Experience = choose_string(origApp.Title_Experience, newApp.Title_Experience)
	merged.Field_Experience = choose_string(origApp.Field_Experience, newApp.Field_Experience)
	merged.Prof_Pic_Url = choose_string(origApp.Prof_Pic_Url, newApp.Prof_Pic_Url)

	merged.Age = choose_int(origApp.Age, newApp.Age)
	merged.Dob = choose_time(origApp.Dob, newApp.Dob)

	return merged
}
