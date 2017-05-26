package models

import (
	"time"
	"database/sql"
	"log"
)

type my_error struct {
	err string
}

func (e *my_error) Error() string {
	return e.err
}

type Applicant struct {
	Name string                `json:"name"`
	ID   int                `json:"id"`

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

func Get_applicant(db *sql.DB, id int) (*Applicant, error) {
	app := new(Applicant)
	e := db.QueryRow("SELECT * FROM applicants WHERE id=?", id).Scan(&app.ID, &app.Name, &app.Email, &app.Password, &app.Dob, &app.Age, &app.Bio, &app.City, &app.State, &app.Title, &app.Field, &app.Title_Experience, &app.Field_Experience, &app.Prof_Pic_Url)

	if e != nil {
		return nil, e
	}

	return app, nil
}

func Store_applicant(db *sql.DB, applicant *Applicant, withID bool) (*Applicant, error) {
	query := insert_query(*applicant, "applicant", withID)
	res, err := db.Exec(query)

	if err != nil {
		return nil, err
	}

	id, _ := res.LastInsertId()
	log.Println("Applicant was stored with id:", id)

	return Get_applicant(db, int(id))
}

func Update_applicant(db *sql.DB, applicant *Applicant) (*Applicant, error) {
	if applicant.ID <= 0 {
		return nil, &my_error{"Send me a valid id"}
	}
	if !applicant_exists(db, applicant.ID) {
		return nil, &my_error{"Applicant doesn't exist"}
	}

	storedApp, err := Get_applicant(db, applicant.ID)
	if err != nil {
		return nil, err
	}

	mergedApp := merge_applicants(storedApp, applicant)

	err = Delete_applicant(db, mergedApp.ID)
	if err != nil {
		return nil, err
	}

	return Store_applicant(db, mergedApp, true)
}

func Delete_applicant(db *sql.DB, id int) error {
	res, err := db.Exec("DELETE FROM applicants WHERE id=?", id)
	insert_id, _ := res.LastInsertId()
	log.Printf("Applicant with id: %d was deleted", insert_id)

	return err
}
