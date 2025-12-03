resource "cloudflare_pages_project" "%[1]s" {
	account_id = "%[2]s"
	name = "%[3]s"
	production_branch = "main"
	deployment_configs = {
		preview = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			env_vars = {
				VAR_ONE = { type = "plain_text", value = "value1" }
			}
		}
		production = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			env_vars = {}
		}
	}
}

