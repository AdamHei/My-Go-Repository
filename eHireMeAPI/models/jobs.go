package models

type Job struct {
	ID int `json:"id"`
	Employer_ID int `json:"employer_id"`
	Description string `json:"description"`
	Compensation string `json:"compensation"`
	Title string `json:"title"`
	Field string `json:"field"`
	Title_experience string `json:"title_experience"`
	Field_experience string `json:"field_experience"`
	City string `json:"city"`
	State string `json:"state"`
	Active bool `json:"active"`
}

//TODO All jobs


//TODO Job by ID


//TODO Jobs by employer id


//TODO Jobs an applicant has matched with


//TODO Applicants that have applied towards a job


//TODO Create


//TODO Update


//TODO Delete


//TODO Search


//TODO Jobs sorted for applicant