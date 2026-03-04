resource "cloudflare_snippet" "%[2]s" {
  zone_id     = "%[1]s"
  name        = "%[2]s"
  main_module = "main.js"

  files {
    name    = "main.js"
    content = "export default {async fetch(request) {return fetch(request)}};"
  }
}