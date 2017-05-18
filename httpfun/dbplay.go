package main

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

//type Applicant struct {
//	id, age                                                                                                  int
//	name, email, password, dob, bio, city, state, title, field, title_experience, field_experience, prof_pic string
//}

func main() {
	db, err := sql.Open("mysql", "root:Spyrohurricane17@/eHireMe")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	rows, e := db.Query("SELECT * FROM applicants")
	if e != nil {
		log.Fatal(e)
	}

	var app Applicant

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&app.name, &app.id, &app.email, &app.password, &app.dob, &app.age, &app.bio, &app.city, &app.state, &app.title, &app.field, &app.title_experience, &app.field_experience, &app.prof_pic)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(app)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
