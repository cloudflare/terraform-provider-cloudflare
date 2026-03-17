resource "cloudflare_pages_project" "%s" {
  account_id        = "%s"
  name              = "%s"
  production_branch = "main"

  deployment_configs = {
    preview = {
      compatibility_date = "2024-01-01"
    }
    production = {
      compatibility_date = "2024-01-01"
    }
  }
}
