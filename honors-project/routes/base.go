package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/adamhei/honors-project/handlers"
)

// HTTP methods
const (
	GET = "GET"
)

type route struct {
	Name, Method, Pattern string
	HandlerFunc           http.HandlerFunc
}

type RouteList []route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods(GET).
		Path("/").
		Name("Index").
		HandlerFunc(handlers.Index)

	router.Methods(GET).
		Path("/exchange-data/all").
		Name("AllExchangeData").
		HandlerFunc(handlers.AllBidAskData)

	router.Methods(GET).
		Path("/exchange-data/biggest-spread").
		Name("BiggestSpread").
		HandlerFunc(handlers.BiggestSpread)

	router.Headers("Content-Type", "application/json")
	return router
}