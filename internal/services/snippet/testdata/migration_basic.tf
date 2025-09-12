resource "cloudflare_snippet" "%[1]s" {
  zone_id     = "%[2]s"
  name        = "%[1]s"
  main_module = "main.js"
  files {
    name    = "main.js"
    content = "export default {async fetch(request) {return fetch(request)}};"
  }
}