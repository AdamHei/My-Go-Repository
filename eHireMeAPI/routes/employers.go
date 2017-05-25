package routes

import "github.com/adamhei/eHireMeAPI/handlers"

func get_employer_routes(env *handlers.Env) Route_list {
	return Route_list{
		route{
			"All_employers",
			"GET",
			"/employers/all/",
			env.All_employers,
		},
		route{
			"EmployerById",
			"GET",
			"/employers/id/{id}",
			env.Employer_id,
		},
		route{
			"Create_employer",
			"POST",
			"/employers/create/",
			env.Create_employer,
		},
		route{
			"Update_employer",
			"POST",
			"/employers/update/",
			env.Update_employer,
		},
	}
}
