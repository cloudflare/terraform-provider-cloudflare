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
}

# Then query it with the data source
data "cloudflare_snippet_rules_list" "%[1]s" {
  zone_id = "%[2]s"
  depends_on = [cloudflare_snippet_rules.%[1]s]
}