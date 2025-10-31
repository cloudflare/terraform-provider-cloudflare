resource "cloudflare_worker" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"
}

resource "cloudflare_worker_version" "%[1]s" {
  account_id = "%[2]s"
  worker_id = cloudflare_worker.%[1]s.id
  modules = [
    {
      name           = "index.js"
      content_base64 = base64encode("export default {async fetch() {return new Response('Updated from base64!')}}")
      content_type   = "application/javascript+module"
    }
  ]
  main_module = "index.js"
}

