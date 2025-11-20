resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "HTTP quarantine policy"
  precedence  = 12402
  action      = "quarantine"
  filters     = ["http"]
  traffic     = "any(http.request.uri.content_category[*] in {35})"

  rule_settings = {
    quarantine = {
      file_types = ["pdf"]
    }
  }
}