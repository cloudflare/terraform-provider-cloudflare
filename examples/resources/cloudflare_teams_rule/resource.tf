resource "cloudflare_teams_rule" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "office"
  description = "desc"
  precedence  = 1
  action      = "block"
  filters     = ["http"]
  traffic     = "http.request.uri == \"https://www.example.com/malicious\""
  rule_settings {
    block_page_enabled = true
    block_page_reason  = "access not permitted"
  }
}
