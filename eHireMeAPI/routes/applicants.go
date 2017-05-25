package routes

import "github.com/adamhei/eHireMeAPI/handlers"

func get_applicant_routes(env *handlers.Env) Route_list {
	return Route_list{
		route{
			"All_applicants",
			"GET",
			"/applicants/all/",
			env.All_applicants,
		},
		route{
			"ApplicantById",
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
}
