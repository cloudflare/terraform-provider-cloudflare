resource "cloudflare_pages_project" "%s" {
  account_id        = "%s"
  name              = "%s"
  production_branch = "main"

  deployment_configs {
    preview {
      compatibility_date = "2024-01-01"
    }
    production {
      environment_variables = {
        NODE_ENV = "production"
        API_URL  = "https://api.example.com"
      }
      secrets = {
        API_KEY     = "secret123"
        DB_PASSWORD = "dbpass456"
      }
      placement {
        mode = "smart"
      }
    }
  }
}
