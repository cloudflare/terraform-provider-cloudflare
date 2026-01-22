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
        SECRET_VAR = {
          type = "secret_text"
          value = "secret-preview-value"
        }
      }
      kv_namespaces = {
        KV_BINDING = {
          namespace_id = cloudflare_workers_kv_namespace.%[1]s_kv_namespace.id
        }
      }
      d1_databases = {
        D1_BINDING = {
          id = "445e2955-951a-4358-a35b-a4d0c813f63"
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
        PROD_SECRET = {
          type = "secret_text"
          value = "secret-prod-value"
        }
      }
      kv_namespaces = {
        KV_BINDING = {
          namespace_id = cloudflare_workers_kv_namespace.%[1]s_kv_namespace.id
        }
      }
      compatibility_date = "2022-08-15"
    }
  }
}

