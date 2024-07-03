
		resource "cloudflare_pages_project" "%[1]s" {
		  account_id = "%[2]s"
		  name = "%[1]s"
		  production_branch = "main"
		  build_config = {
  build_caching = true
			build_command = "npm run build"
			destination_dir = "build"
			root_dir = "/"
			web_analytics_tag = "cee1c73f6e4743d0b5e6bb1a0bcaabcc"
			web_analytics_token = "021e1057c18547eca7b79f2516f06o7x"
}
		}
		