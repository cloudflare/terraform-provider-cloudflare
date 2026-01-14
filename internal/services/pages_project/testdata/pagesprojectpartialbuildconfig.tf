resource "cloudflare_pages_project" "%[1]s" {
	account_id = "%[2]s"
	name = "%[3]s"
	production_branch = "main"

	# Complete build_config - specifying all fields including build_caching
	# This tests that after import, specifying build_config with all fields works
	build_config = {
		build_caching = false
		build_command = "yarn build"
		destination_dir = "dist"
		root_dir = "/"
	}

	deployment_configs = {
		preview = {
			compatibility_date = "2023-08-25"
		}
		production = {
			compatibility_date = "2023-08-25"
		}
	}
}

