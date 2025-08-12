resource "cloudflare_snippet" "%[1]s" {
  zone_id      = "%[2]s"
  snippet_name = "rules_set_snippet"
  files = [
    {
      name    = "main.js"
      content = <<-EOT
      export default {
        async fetch(request) {
          return fetch(request);
        },
      }
      EOT
    }
  ]
  metadata = {
    main_module = "main.js"
  }
}

# First create a snippet rule to query
resource "cloudflare_snippet_rules" "%[1]s" {
  zone_id = "%[2]s"
  rules = [
    {
      snippet_name = "rules_set_snippet"
      expression   = "http.request.uri.path contains \"/datasource-test\""
      enabled      = true
      description  = "Data source test snippet rule"
    }
  ]
  depends_on = [cloudflare_snippet.%[1]s]
}

# Then query it with the data source
data "cloudflare_snippet_rules_list" "%[1]s" {
  zone_id = "%[2]s"
  depends_on = [cloudflare_snippet_rules.%[1]s]
}