package models

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Employer struct {
	ID           int `json:"id"`
	Company      string `json:"company"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Description  string `json:"description"`
	Prof_Pic_Url string `json:"prof_pic_url"`
}

// Get_employer returns an employer by ID from the database
func Get_employer(db *sql.DB, id int) (*Employer, error) {
	emp := new(Employer)
	e := db.QueryRow("SELECT * FROM employers WHERE id=?", id).Scan(&emp.ID,
		&emp.Company, &emp.Password, &emp.Email, &emp.Description, &emp.Prof_Pic_Url)

	if e != nil {
		return nil, e
	}
	return emp, nil
}

// All_employers returns every employer in the corresponding table
func All_employers(db *sql.DB) ([]*Employer, error) {
	rows, err := db.Query("SELECT * FROM employers")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	employers := make([]*Employer, 0)
	for rows.Next() {
		emp := new(Employer)
		err := rows.Scan(&emp.ID,
			&emp.Company, &emp.Password, &emp.Email, &emp.Description, &emp.Prof_Pic_Url)
		if err != nil {
			return nil, err
		}
		employers = append(employers, emp)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return employers, nil
}

// Store_employer will, given an employer, attempt to insert it into the proper table
// withID = false when inserting the first time
// true when performing a partial update
func Store_employer(db *sql.DB, employer *Employer) (*Employer, error) {
	res, err := insert_model(employer, emp_insert_statement, db)

	if err != nil {
		//Couldn't insert
		return nil, err
	}

	id, _ := res.LastInsertId()
	fmt.Println("Employer was stored with id:", id)

	return Get_employer(db, int(id))
}

// Update_employer will attempt to perform a partial update by merging employers and returning the new form
func Update_employer(db *sql.DB, employer *Employer) (*Employer, error) {
	if employer.ID <= 0 {
		return nil, &My_error{"Send a valid id", http.StatusBadRequest}
	}
	if !employer_exists(db, employer.ID) {
		return nil, &My_error{"Employer doesn't exist", http.StatusBadRequest}
	}

	stored_emp, err := Get_employer(db, employer.ID)
	if err != nil {
		return nil, err
	}

	merged_emp := merge_employers(stored_emp, employer)

	_, err = update_model(merged_emp, emp_update_statement, db)

	if err != nil {
		return nil, err
	}

	return Get_employer(db, merged_emp.ID)
}

// Delete_employer will remove an employer from the database by id
func Delete_employer(db *sql.DB, id int) error {
	res, err := db.Exec("DELETE FROM employers WHERE id=?", id)
	fmt.Println(res.LastInsertId())

	return err
}