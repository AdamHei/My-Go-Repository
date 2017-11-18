package routes

import (
	"github.com/adamhei/honors-project/handlers"
	"github.com/gorilla/mux"
	"net/http"
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

	router.Methods(GET).
		Path("/historical-data/compare/{exchange1}/{exchange2}").
		Name("CompareHistory").
		HandlerFunc(handlers.Compare)

	router.Headers("Content-Type", "application/json")
	return router
}
