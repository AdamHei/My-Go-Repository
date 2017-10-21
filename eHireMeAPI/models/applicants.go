package models

import (
	"time"
	"database/sql"
	"log"
	"net/http"
)

// type Applicant is the struct representation of Applicants in the database, with corresponding JSON keys
type Applicant struct {
	Name             string    `json:"name"`
	ID               int       `json:"id"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	Dob              time.Time `json:"dob"`
	Age              int       `json:"age"`
	Bio              string    `json:"bio"`
	City             string    `json:"city"`
	State            string    `json:"state"`
	Title            string    `json:"title"`
	Field            string    `json:"field"`
	Title_Experience string    `json:"title_experience"`
	Field_Experience string    `json:"field_experience"`
	Prof_Pic_Url     string    `json:"prof_pic_url"`
}

// All_applicants will, given a database, attempt to return all applicants in the corresponding table
func All_applicants(db *sql.DB) ([]*Applicant, error) {
	rows, err := db.Query("SELECT * FROM applicants")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	apps := make([]*Applicant, 0)
	for rows.Next() {
		app := new(Applicant)
		err = rows.Scan(&app.ID, &app.Name, &app.Email, &app.Password, &app.Dob, &app.Age, &app.Bio, &app.City, &app.State, &app.Title, &app.Field, &app.Title_Experience, &app.Field_Experience, &app.Prof_Pic_Url)
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

// Get_applicant will attempt to return an applicant by id
func Get_applicant(db *sql.DB, id int) (*Applicant, error) {
	app := new(Applicant)
	e := db.QueryRow("SELECT * FROM applicants WHERE id=?", id).Scan(&app.ID, &app.Name, &app.Email,
		&app.Password, &app.Dob, &app.Age, &app.Bio, &app.City, &app.State, &app.Title, &app.Field,
		&app.Title_Experience, &app.Field_Experience, &app.Prof_Pic_Url)

	if e != nil {
		return nil, e
	}

	return app, nil
}

// Store_applicant will attempt to insert a given applicant in the proper table
// withID = false when inserting for the first time
// true when updating an applicant
func Store_applicant(db *sql.DB, applicant *Applicant) (*Applicant, error) {
	res, err := insert_model(applicant, app_insert_statement, db)

	if err != nil {
		return nil, err
	}

	id, _ := res.LastInsertId()
	log.Println("Applicant was stored with id:", id)

	return Get_applicant(db, int(id))
}

// Update_applicant will attempt to perform a partial update given an applicant and return the new applicant
func Update_applicant(db *sql.DB, applicant *Applicant) (*Applicant, error) {
	if applicant.ID <= 0 {
		return nil, &My_error{"Send me a valid id", http.StatusBadRequest}
	}
	if !applicant_exists(db, applicant.ID) {
		return nil, &My_error{"Applicant doesn't exist", http.StatusBadRequest}
	}

	storedApp, err := Get_applicant(db, applicant.ID)
	if err != nil {
		return nil, err
	}

	mergedApp := merge_applicants(storedApp, applicant)
	_, err = update_model(mergedApp, app_update_statement, db)

	if err != nil {
		return nil, err
	}

	return Get_applicant(db, mergedApp.ID)
}

// Delete_applicant will attempt to delete an applicant by id and return an error if any
func Delete_applicant(db *sql.DB, id int) error {
	res, err := db.Exec("DELETE FROM applicants WHERE id=?", id)
	insert_id, _ := res.LastInsertId()
	log.Printf("Applicant with id: %d was deleted", insert_id)

	return err
}
