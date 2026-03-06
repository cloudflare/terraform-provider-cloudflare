resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"

  deployment_configs = {
    preview = {
      always_use_latest_compatibility_date = true
    }
    production = {
      always_use_latest_compatibility_date = false
    }
  }
}
