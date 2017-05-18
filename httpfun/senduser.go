package main

import (
	"net/http"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
)

type Applicant struct {
	id, age                                                                                                  int
	name, email, password, dob, bio, city, state, title, field, title_experience, field_experience, prof_pic string
}

func getApplicant(id string) Applicant {
	db, err := sql.Open("mysql", "root:Spyrohurricane17@/eHireMe")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var app Applicant

	e := db.QueryRow("SELECT name, email FROM applicants WHERE id=?", id).Scan(&app.name, &app.email)
	if e != nil {
		fmt.Println(e)
	}

	return app
}

func applicant(writer http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Path[len("/applicant/"):]
	if len(user_id) == 0 {
		fmt.Fprint(writer, "Give me an id")
		return
	}

	app := getApplicant(user_id)

	if app.name != "" {
		fmt.Fprintf(writer, "You selected user %s with email: %s", app.name, app.email)
	} else {
		fmt.Fprint(writer, "We couldn't find that user!")
	}
}

func Index(writer http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(writer, "You found my webpage!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
