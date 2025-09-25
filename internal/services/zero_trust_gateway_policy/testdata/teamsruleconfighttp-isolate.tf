resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  name = "%[1]s"
  account_id = "%[2]s"
  description = "desc"
  precedence = 12307
  action = "isolate"
  enabled = true
  filters = ["http"]
  traffic = "any(http.request.uri.security_category[*] in {21}) or any(http.request.uri.content_category[*] in {32})"
  rule_settings = {
    biso_admin_controls = {
      copy = "remote_only"
      download = "enabled"
      keyboard = "enabled"
      paste = "enabled"
      printing = "enabled"
      upload = "enabled"
      version = "v1"

    }
  }
}
