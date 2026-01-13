resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "Response Firewall Ruleset %[2]s"
  phase   = "http_response_firewall_managed"
  kind    = "zone"

  rules {
    expression = "http.response.code eq 403"
    action     = "log"
    description = "Log 403 responses"
  }

  rules {
    expression = "http.response.headers[\"x-custom-header\"][0] eq \"sensitive\""
    action     = "log"
    description = "Log responses with sensitive header"
  }
}
