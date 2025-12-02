resource "cloudflare_pages_project" "%[1]s" {
	account_id = "%[2]s"
	name = "%[3]s"
	production_branch = "main"
	
	deployment_configs = {
		preview = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			env_vars = {
				PLAIN_TEXT_VAR = {
					type = "plain_text"
					value = "plain-text-value"
				}
				SECRET_VAR = {
					type = "secret_text" 
					value = "secret-value-123"
				}
			}
		}
	
		production = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			env_vars = {
				PROD_PLAIN = {
					type = "plain_text"
					value = "production-plain"
				}
				PROD_SECRET = {
					type = "secret_text"
					value = "production-secret-456"
				}
			}
		}
	}
}
