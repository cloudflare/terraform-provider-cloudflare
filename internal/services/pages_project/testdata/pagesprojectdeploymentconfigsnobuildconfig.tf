resource "cloudflare_pages_project" "%[1]s" {
	account_id = "%[2]s"
	name = "%[3]s"
	production_branch = "main"

	deployment_configs = {
		preview = {
			always_use_latest_compatibility_date = false
			compatibility_date = "2023-08-25"
			fail_open = false
			usage_model = "bundled"
		}
		production = {
			always_use_latest_compatibility_date = false
			compatibility_date = "2023-08-25"
			fail_open = false
			usage_model = "bundled"
		}
	}
}

