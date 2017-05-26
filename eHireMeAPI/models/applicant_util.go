package models

import (
	"database/sql"
	"fmt"
)

func (app Applicant) member_fields(withID bool) string {
	fields := "name, email, password, dob, age, bio, city, state, title, field, title_experience, field_experience, prof_pic)"
	if withID {
		fields = "(id, " + fields
	} else {
		fields = "(" + fields
	}
	return fields
}

func (app Applicant) member_values(withID bool) string {
	dob := app.Dob.Format("YYY-MM-DD")
	values := fmt.Sprintf("%s, %s, %s, %s, %d, %s, %s, %s, %s, %s, %s, %s, %s)",
		app.Name, app.Email, app.Password, dob, app.Age, app.Bio,
		app.City, app.State, app.Title, app.Field, app.Title_Experience,
		app.Field_Experience, app.Prof_Pic_Url)
	if withID {
		return add_ID(values, app.ID)
	} else {
		return "(" + values
	}
}

func applicant_exists(db *sql.DB, query_id int) bool {
	id := -1
	err := db.QueryRow("SELECT id FROM applicants WHERE id=?", query_id).Scan(&id)

	return id != -1 || err == nil
}

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
