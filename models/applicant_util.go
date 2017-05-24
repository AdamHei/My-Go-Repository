package models

import (
	"database/sql"
	"time"
)

func applicant_exists(db *sql.DB, query_id int) bool {
	id := -1
	err := db.QueryRow("SELECT id FROM applicants WHERE id=?", query_id).Scan(&id)

	return id != -1 || err == nil
}

func choose_string(att1, att2 string) string {
	if att2 == "" || att1 == att2 {
		return att1
	}
	return att2
}

func choose_int(i1, i2 int) int {
	if i2 == 0 || i1 == i2 {
		return i1
	}
	return i2
}

func choose_time(t1, t2 time.Time) time.Time {
	if t2.IsZero() || t1.Equal(t2) {
		return t1
	}
	return t2
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
