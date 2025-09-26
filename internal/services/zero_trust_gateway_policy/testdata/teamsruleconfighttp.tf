resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  name = "%[1]s"
  account_id = "%[2]s"
  description = "desc"
  precedence = 12305
  action = "allow"
  filters = ["http"]
  traffic = "any(http.request.uri.security_category[*] in {22}) or any(http.request.uri.content_category[*] in {34})"
  rule_settings = {
    add_headers = {"Xhello": ["abcd", "efg"]}
    check_session = {
      duration = "1h2m9s"
      enforce = true
    }
  }
}
