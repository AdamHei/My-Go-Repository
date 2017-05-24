package models

import (
	"database/sql"
	"time"
)

func chooseString(att1, att2 string) string {
	if att2 == "" || att1 == att2 {
		return att1
	}
	return att2
}

func chooseInt(i1, i2 int) int {
	if i2 == 0 || i1 == i2 {
		return i1
	}
	return i2
}

func chooseTime(t1, t2 time.Time) time.Time {
	if t2.IsZero() || t1.Equal(t2) {
		return t1
	}
	return t2
}

func mergeApplicants(origApp *Applicant, newApp *Applicant) *Applicant {
	merged := new(Applicant)
	//TODO Change password in a diff function
	//ID is immutable; Create new user
	merged.Id = origApp.Id
	merged.Password = origApp.Password
	merged.Name = chooseString(origApp.Name, newApp.Name)
	merged.Email = chooseString(origApp.Email, newApp.Email)
	merged.Bio = chooseString(origApp.Bio, newApp.Bio)
	merged.City = chooseString(origApp.City, newApp.City)
	merged.State = chooseString(origApp.State, newApp.State)
	merged.Title = chooseString(origApp.Title, newApp.Title)
	merged.Field = chooseString(origApp.Field, newApp.Field)
	merged.Title_Experience = chooseString(origApp.Title_Experience, newApp.Title_Experience)
	merged.Field_Experience = chooseString(origApp.Field_Experience, newApp.Field_Experience)
	merged.Prof_Pic_Url = chooseString(origApp.Prof_Pic_Url, newApp.Prof_Pic_Url)

	merged.Age = chooseInt(origApp.Age, newApp.Age)
	merged.Dob = chooseTime(origApp.Dob, newApp.Dob)

	return merged
}

func applicantExists(db *sql.DB, applicant *Applicant) bool {
	id := -1
	err := db.QueryRow("SELECT id FROM applicants WHERE id=?", applicant.Id).Scan(&id)

	//TODO May need to return error as well
	return id != -1 || err == nil
}
