package main

import (
	"github.com/adamhei/hackisuhonors/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
