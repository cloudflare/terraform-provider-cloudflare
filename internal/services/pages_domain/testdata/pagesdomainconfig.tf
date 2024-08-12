
		resource "cloudflare_pages_project" "%[1]s" {
			account_id = "%[2]s"
			name = "%[3]s"
			production_branch = "main"
		}
		resource "cloudflare_pages_domain" "%[1]s" {
		  account_id = "%[2]s"
		  project_name = cloudflare_pages_project.%[1]s.name
		  domain = "%[4]s"
		}
		