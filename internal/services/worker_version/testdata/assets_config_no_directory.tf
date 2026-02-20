resource "cloudflare_worker" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"
}

resource "cloudflare_worker_version" "%[1]s" {
  account_id  = "%[2]s"
  worker_id   = cloudflare_worker.%[1]s.id
  main_module = "index.js"

  modules = [{
    name         = "index.js"
    content_file = "%[3]s"
    content_type = "application/javascript+module"
  }]

  assets = {
    config = {
      not_found_handling = "single-page-application"
      run_worker_first   = ["/api/v1/*"]
    }
  }
}
