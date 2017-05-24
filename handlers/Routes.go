package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name, Method, Pattern string
	HandlerFunc           http.HandlerFunc
}

type Routes []Route

func NewRouter(env *Env) *mux.Router {
	var routes = Routes{
		Route{
			"Index",
			"GET",
			"/",
			env.Index,
		},
		Route{
			"AllApplicants",
			"GET",
			"/applicants/all/",
			env.AllApplicants,
		},
		Route{
			"ById",
			"GET",
			"/applicants/id/{id}",
			env.ApplicantId,
		},
		Route{
			"CreateApplicant",
			"POST",
			"/applicants/create/",
			env.CreateApplicant,
		},
		Route{
			"UpdateApplicant",
			"POST",
			"/applicants/update/",
			env.UpdateApplicant,
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
