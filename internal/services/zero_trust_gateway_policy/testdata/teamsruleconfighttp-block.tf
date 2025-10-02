resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  name = "%[1]s"
  account_id = "%[2]s"
  description = "desc"
  precedence = 12306
  action = "block"
  filters = ["http"]
  traffic = "any(http.request.uri.security_category[*] in {21}) or any(http.request.uri.content_category[*] in {32})"
  rule_settings = {
    block_page = {
      target_uri = "https://examples.com"
      include_context = false
    }
    notification_settings = {"enabled": true, "msg": "msg"}
  }
}
