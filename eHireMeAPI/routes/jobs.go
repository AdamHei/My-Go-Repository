package routes

import "github.com/adamhei/eHireMeAPI/handlers"

// get_jobs_routes returns a Route_list as a registry of all Jobs routes
func get_jobs_routes(env *handlers.Env) Route_list {
	return Route_list{
		route{
			"All_jobs",
			GET,
			"/jobs/all/",
			env.All_jobs,
		},
		route{
			"Job_By_ID",
			GET,
			"/jobs/id/{id}",
			env.Job_id,
		},
		route{
			"Job_By_Employer",
			GET,
			"/jobs/employer/{id}",
			env.Jobs_by_employer,
		},
		route{
			"Create_job",
			POST,
			"/jobs/create/",
			env.Create_job,
		},
		route{
			"Update_job",
			POST,
			"/jobs/update/",
			env.Update_job,
		},
		route{
			"Delete_job",
			DELETE,
			"/jobs/delete/{id}",
			env.Delete_job,
		},
	}
}
