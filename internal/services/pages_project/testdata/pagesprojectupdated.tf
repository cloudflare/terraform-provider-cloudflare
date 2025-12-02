resource "cloudflare_pages_project" "%[1]s" {
	account_id = "%[2]s"
	name = "%[3]s"
	production_branch = "develop"
	
	build_config = {
		build_caching = true
		build_command = "yarn build"
		destination_dir = "build"
		root_dir = "/"
	}

	deployment_configs = {
		preview = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			
			env_vars = {
				UPDATED_VAR = {
					type = "plain_text"
					value = "updated-value"
				}
			}
		}
	
		production = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			
			env_vars = {
				PROD_UPDATED = {
					type = "secret_text"
					value = "updated-secret"
				}
			}
		}
	}
}
