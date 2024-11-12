
		resource "cloudflare_pages_project" "%[1]s" {
		  account_id = "%[2]s"
		  name = "%[1]s"
		  production_branch = "main"
		  source = [{
			type = "github"
			config = [{
				owner = "%[4]s"
				repo_name = "%[5]s"
				production_branch = "main"
				pr_comments_enabled = true
				deployments_enabled = true
				production_deployment_enabled = true
				preview_deployment_setting = "custom"
				preview_branch_includes = ["dev","preview"]
				preview_branch_excludes = ["main", "prod"]
			}]
		  }]
		}
		