resource "cloudflare_worker" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"
}

resource "cloudflare_worker_version" "%[1]s" {
  account_id = "%[2]s"
  worker_id = cloudflare_worker.%[1]s.id
  modules = [
    {
      name         = "index.js"
      content_file = "%[5]s"
      content_type = "application/javascript+module"
    }
  ]
  main_module = "index.js"
  assets = {
    directory = "%[3]s"
    config = {
      run_worker_first = %[4]s
    }
  }
}

