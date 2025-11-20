resource "cloudflare_pages_project" "%[1]s" {
	account_id = "%[2]s"
	name = "%[3]s"
	production_branch = "main"
	
	source = {
		type = "github"
		config = {
			owner = "%[4]s"
			repo_name = "%[5]s"
			production_branch = "main"
			preview_deployment_setting = "%[6]s"
			%[7]s
		}
	}
}
