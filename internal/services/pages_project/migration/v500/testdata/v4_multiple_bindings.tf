resource "cloudflare_pages_project" "%s" {
  account_id        = "%s"
  name              = "%s"
  production_branch = "main"

  build_config {
    build_caching   = false
    build_command   = ""
    destination_dir = ""
    root_dir        = ""
  }

  deployment_configs {
    preview {
      compatibility_date = "2024-01-01"
    }
    production {
      compatibility_date = "2024-01-01"
      kv_namespaces = {
        KV_BINDING_1 = "%s"
        KV_BINDING_2 = "%s"
      }
      d1_databases = {
        DB_BINDING_1 = "%s"
        DB_BINDING_2 = "%s"
      }
      r2_buckets = {
        BUCKET_1 = "my-bucket-1"
        BUCKET_2 = "my-bucket-2"
      }
      durable_object_namespaces = {
        DO_BINDING_1 = "%s"
        DO_BINDING_2 = "%s"
      }
    }
  }
}
