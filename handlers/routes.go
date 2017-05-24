package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
)

type route struct {
	Name, Method, Pattern string
	HandlerFunc           http.HandlerFunc
}

type routes []route

func NewRouter(env *Env) *mux.Router {
	var routes = routes{
		route{
			"Index",
			"GET",
			"/",
			env.Index,
		},
		route{
			"all_applicants",
			"GET",
			"/applicants/all/",
			env.All_applicants,
		},
		route{
			"ById",
			"GET",
			"/applicants/id/{id}",
			env.Applicant_id,
		},
		route{
			"Create_applicant",
			"POST",
			"/applicants/create/",
			env.Create_applicant,
		},
		route{
			"Update_applicant",
			"POST",
			"/applicants/update/",
			env.Update_applicant,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	router.Headers("Content-Type", "application/json")

	return router
}
