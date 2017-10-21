package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/adamhei/hackisuhonors/handlers"
)

// HTTP methods
const (
	GET    = "GET"
	POST   = "POST"
)

type route struct {
	Name, Method, Pattern string
	HandlerFunc           http.HandlerFunc
}

type RouteList []route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods(GET).
		Path("/exchange-data/all").
		Name("AllExchangeData").
		HandlerFunc(handlers.AllBidAskData)

	router.Headers("Content-Type", "application/json")
	return router
}
