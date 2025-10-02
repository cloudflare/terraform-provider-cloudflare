resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "HTTP redirect policy"
  precedence  = 12401
  action      = "redirect"
  filters     = ["http"]
  traffic     = "any(http.request.uri.security_category[*] in {25})"

  rule_settings = {
    redirect = {
      target_uri                = "https://redirect.example.com"
      include_context           = true
      preserve_path_and_query   = true
    }
  }
}