resource "cloudflare_workers_kv_namespace" "%[1]s_kv_namespace" {
  account_id = "%[2]s"
  title = "tfacctest-pages-project-kv-namespace-%[1]s"
}

resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"
  deployment_configs = {
    preview = {
      env_vars = {
        ENVIRONMENT = {
          type = "plain_text"
          value = "preview"
        }
      }
      compatibility_date = "2022-08-15"
    }
    production = {
      env_vars = {
        ENVIRONMENT = {
          type = "plain_text"
          value = "production"
        }
      }
      compatibility_date = "2022-08-15"
    }
  }
}

