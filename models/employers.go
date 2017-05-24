package models

import (
	"database/sql"
	"fmt"
)

type Employer struct {
	ID           int `json:"id"`
	Company      string `json:"company"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Description  string `json:"description"`
	Prof_Pic_Url string `json:"prof_pic_url"`
}

func Get_employer(db *sql.DB, id int) (*Employer, error) {
	emp := new(Employer)
	e := db.QueryRow("SELECT * FROM employers WHERE id=?", id).Scan(&emp.ID,
		&emp.Company, &emp.Password, &emp.Email, &emp.Description, &emp.Prof_Pic_Url)

	if e != nil {
		return nil, e
	}
	return emp, nil
}

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

func Store_employer(db *sql.DB, employer *Employer) (*Employer, error) {
	res, err := db.Exec("INSERT INTO employers (company, password, email, description, prof_pic) "+
		"VALUES (?, ?, ?, ?, ?)",
		employer.Company, employer.Password, employer.Email, employer.Description, employer.Prof_Pic_Url)

	if err != nil {
		//Couldn't insert
		return nil, err
	}

	id, _ := res.LastInsertId()
	fmt.Println("Employer was stored with id:", id)

	return Get_employer(db, int(id))
}

func Delete_employer(db *sql.DB, id int) error {
	res, err := db.Exec("DELETE FROM employers WHERE id=?", id)
	fmt.Println(res.LastInsertId())

	return err
}

func Update_employer(db *sql.DB, employer *Employer) (*Employer, error) {
	if employer.ID <= 0 {
		return nil, &my_error{"Send a valid id"}
	}
	if !employer_exists(db, employer.ID) {
		return nil, &my_error{"Employer doesn't exist"}
	}

	stored_emp, err := Get_employer(db, employer.ID)
	if err != nil {
		return nil, err
	}

	merged_emp := merge_employers(stored_emp, employer)

	err = Delete_employer(db, merged_emp.ID)
	if err != nil {
		return nil, err
	}

	return Store_employer(db, merged_emp)
}
