resource "cloudflare_snippet" "%[1]s" {
  zone_id     = "%[2]s"
  name        = "rules_set_snippet"
  main_module = "main.js"

  files {
    name    = "main.js"
    content = "export default {async fetch(request) {return fetch(request)}};"
  }
}

resource "cloudflare_snippet_rules" "%[1]s" {
  zone_id = "%[2]s"

  rules {
    snippet_name = "rules_set_snippet"
    expression   = "http.request.uri.path contains \"/v1\""
    enabled      = true
    description  = "First rule"
  }

  rules {
    snippet_name = "rules_set_snippet"
    expression   = "http.request.uri.path contains \"/v2\""
    enabled      = false
    description  = "Second rule (disabled)"
  }

  depends_on = [cloudflare_snippet.%[1]s]
}