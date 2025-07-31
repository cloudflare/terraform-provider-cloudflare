resource "cloudflare_snippet_rules" "%[1]s" {
  zone_id = "%[2]s"
  rules = [
    {
      snippet_name = "rules_set_snippet"
      expression   = "http.request.uri.path contains \"/updated\""
      enabled      = false
      description  = "Updated test snippet rule"
    }
  ]
}