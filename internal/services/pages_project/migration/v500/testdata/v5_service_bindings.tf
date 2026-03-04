resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"

  build_config = {
    build_caching   = false
    build_command   = ""
    destination_dir = ""
    root_dir        = ""
  }

  deployment_configs = {
    preview = {
      compatibility_date = "2026-01-14"
    }
    production = {
      services = {
        MY_SERVICE_1 = {
          service     = "worker-1"
          environment = "production"
        }
        MY_SERVICE_2 = {
          service = "worker-2"
        }
        MY_SERVICE_3 = {
          service     = "worker-3"
          environment = "staging"
        }
      }
    }
  }
}
