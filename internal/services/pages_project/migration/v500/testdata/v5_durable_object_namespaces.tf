resource "cloudflare_pages_project" "%s" {
  account_id        = "%s"
  name              = "%s"
  production_branch = "main"

  build_config = {
    build_caching   = false
    build_command   = ""
    destination_dir = ""
    root_dir        = ""
  }

  deployment_configs = {
    preview = {
      compatibility_date = "2024-01-01"
    }
    production = {
      compatibility_date = "2024-01-01"
      durable_object_namespaces = {
        MY_DO = "%s"
      }
    }
  }
}
