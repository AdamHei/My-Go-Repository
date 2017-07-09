package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/adamhei/eHireMeAPI/handlers"
)

// HTTP methods
const (
	GET = "GET"
	POST = "POST"
	DELETE = "DELETE"
)

type route struct {
	Name, Method, Pattern string
	HandlerFunc           http.HandlerFunc
}

type Route_list []route

// all_routes returns a Route_list containing the base route and all employer, applicant, and jobs routes
func all_routes(env *handlers.Env) Route_list {
	all_routes := Route_list{
		route{
			"Index",
			GET,
			"/",
			env.Index,
		},
	}

	all_routes = append(all_routes, get_applicant_routes(env)...)
	all_routes = append(all_routes, get_employer_routes(env)...)
	all_routes = append(all_routes, get_jobs_routes(env)...)
	return all_routes
}

// NewRouter returns a new Mux Router with all routes registered and content-type set to JSON
func NewRouter(env *handlers.Env) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range all_routes(env) {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	router.Headers("Content-Type", "application/json")

	return router
}
