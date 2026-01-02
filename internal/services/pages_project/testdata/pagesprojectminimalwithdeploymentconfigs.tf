resource "cloudflare_pages_project" "%[1]s" {
	account_id = "%[2]s"
	name = "%[3]s"
	production_branch = "main"

	build_config = {
		build_caching = true
		build_command = "yarn build"
		destination_dir = "build"
		root_dir = "/"
	}

	deployment_configs = {
		preview = {
			compatibility_date = "2023-06-01"
		}
		production = {
			compatibility_date = "2023-06-01"
		}
	}
}

