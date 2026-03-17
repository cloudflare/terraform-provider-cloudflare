resource "cloudflare_pages_project" "%s" {
  account_id        = "%s"
  name              = "%s"
  production_branch = "main"

  build_config {
    build_command = ""
  }

  deployment_configs {
    preview {
      compatibility_date = "2026-01-14"
    }
    production {
      kv_namespaces = {
        MY_KV = "%s"
      }
      d1_databases = {
        MY_DB = "%s"
      }
      r2_buckets = {
        MY_BUCKET = "%s"
      }
      service_binding {
        name        = "MY_SERVICE"
        service     = "%s"
        environment = "production"
      }
    }
  }
}
