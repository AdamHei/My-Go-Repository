package models

import (
	"database/sql"
	"fmt"
)

// member_fields, part of the Model implementation, returns the column names of the Employer schema, with or without ID
func (employer Employer) member_fields(withID bool) string {
	fields := "company, password, email, description, prof_pic)"
	if withID {
		return "(id, " + fields
	} else {
		return "(" + fields
	}
}

// member_values, part of the Model implementation, returns a formatted string of the member values of a given Employer
// with or without the ID
func (employer Employer) member_values(withID bool) string {
	values := fmt.Sprintf("%s, %s, %s, %s, %s)", employer.Company, employer.Password, employer.Email,
		employer.Description, employer.Prof_Pic_Url)
	if withID {
		return add_ID(values, employer.ID)
	} else {
		return "(" + values
	}
}

func employer_exists(db *sql.DB, query_id int) bool {
	temp_id := -1
	err := db.QueryRow("SELECT id FROM employers WHERE id=?", query_id).Scan(&temp_id)

	return temp_id != -1 || err == nil
}

// merge_employers will, given an employer from the database and one from the client, return the merged version of the two
func merge_employers(original_emp, new_emp *Employer) *Employer {
	merged := new(Employer)
	merged.ID = original_emp.ID
	merged.Password = original_emp.Password

	merged.Company = choose_string(original_emp.Company, new_emp.Company)
	merged.Email = choose_string(original_emp.Email, new_emp.Email)
	merged.Description = choose_string(original_emp.Description, new_emp.Description)
	merged.Prof_Pic_Url = choose_string(original_emp.Prof_Pic_Url, new_emp.Prof_Pic_Url)

	return merged
}
