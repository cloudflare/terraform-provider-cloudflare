		resource "cloudflare_pages_project" "%[1]s" {
		  account_id = "%[2]s"
		  name = "%[1]s"
		  production_branch = "main"
		  
		  source = {
			type = "github"
			config = {
			  owner = "%[3]s"
			  repo_name = "%[4]s"
			  production_branch = "main"
			  preview_deployment_setting = "%[5]s"
			  %[6]s
			}
		  }
		}
		