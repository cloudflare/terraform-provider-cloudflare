resource "cloudflare_pages_project" "%s" {
  account_id        = "%s"
  name              = "%s"
  production_branch = "main"

  deployment_configs = {
    preview = {
      usage_model = "unbound"
    }
    production = {
      usage_model = "standard"
    }
  }
}
