package models

import "database/sql"

func employer_exists(db *sql.DB, query_id int) bool {
	temp_id := -1
	err := db.QueryRow("SELECT id FROM employers WHERE id=?", query_id).Scan(&temp_id)

	return temp_id != -1 || err == nil
}

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
