package models

import (
	"time"
	"database/sql"
	"fmt"
)

type MyError struct {
	err string
}

func (e *MyError) Error() string {
	return e.err
}

type Applicant struct {
	Name string                `json:"name"`
	Id   int                `json:"id"`

	Email    string        `json:"email"`
	Password string        `json:"password"`

	Dob              time.Time                `json:"dob"`
	Age              int               `json:"age"`
	Bio              string        `json:"bio"`
	City             string        `json:"city"`
	State            string        `json:"state"`
	Title            string        `json:"title"`
	Field            string        `json:"field"`
	Title_Experience string        `json:"title_experience"`
	Field_Experience string        `json:"field_experience"`
	Prof_Pic_Url     string        `json:"prof_pic_url"`
}

func AllApplicants(db *sql.DB) ([]*Applicant, error) {
	rows, err := db.Query("SELECT * FROM applicants")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	apps := make([]*Applicant, 0)
	for rows.Next() {
		app := new(Applicant)
		err := rows.Scan(&app.Name, &app.Id, &app.Email, &app.Password, &app.Dob, &app.Age, &app.Bio, &app.City, &app.State, &app.Title, &app.Field, &app.Title_Experience, &app.Field_Experience, &app.Prof_Pic_Url)
		if err != nil {
			return nil, err
		}
		apps = append(apps, app)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return apps, nil
}

func GetApplicant(db *sql.DB, id int) (*Applicant, error) {
	app := new(Applicant)
	e := db.QueryRow("SELECT * FROM applicants WHERE id=?", id).Scan(&app.Name, &app.Id, &app.Email, &app.Password, &app.Dob, &app.Age, &app.Bio, &app.City, &app.State, &app.Title, &app.Field, &app.Title_Experience, &app.Field_Experience, &app.Prof_Pic_Url)

	if e != nil {
		return nil, e
	}

	return app, nil
}

func StoreApplicant(db *sql.DB, applicant *Applicant) error {
	if applicantExists(db, applicant) {
		return &MyError{"Duplicate id"}
	}

	res, err := db.Exec("INSERT INTO applicants "+
		"(name, id, email, password, dob, age, bio, city, state, title, field, title_experience, field_experience, prof_pic)"+
		" VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		applicant.Name, applicant.Id, applicant.Email, applicant.Password, applicant.Dob, applicant.Age, applicant.Bio, applicant.City,
		applicant.State, applicant.Title, applicant.Field, applicant.Title_Experience, applicant.Field_Experience, applicant.Prof_Pic_Url)

	if err != nil {
		//Couldn't insert
		return err
	}

	fmt.Println(res.LastInsertId())
	return nil
}

func DeleteApplicant(db *sql.DB, id int) error {
	res, err := db.Exec("DELETE FROM applicants WHERE id=?", id)
	fmt.Println(res)

	return err
}

func UpdateApplicant(db *sql.DB, applicant *Applicant) error {
	if !applicantExists(db, applicant) {
		return &MyError{"Applicant doesn't exist"}
	}

	storedApp, err := GetApplicant(db, applicant.Id)
	if err != nil {
		return err
	}

	mergedApp := mergeApplicants(storedApp, applicant)

	err = DeleteApplicant(db, mergedApp.Id)
	if err != nil {
		return err
	}

	err = StoreApplicant(db, mergedApp)
	return err
}
