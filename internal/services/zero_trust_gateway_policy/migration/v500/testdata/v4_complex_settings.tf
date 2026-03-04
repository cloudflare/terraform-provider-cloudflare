resource "cloudflare_teams_rule" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-test-complex-%[1]s"
  description = "Policy with complex settings"
  precedence  = 10000
  action      = "allow"
  enabled     = true
  filters     = ["http"]
  traffic     = "http.request.uri matches \".*api.*\""

  rule_settings {
    add_headers  = {}
    override_ips = []

    check_session {
      enforce  = true
      duration = "24h0m0s"
    }

    payload_log {
      enabled = true
    }
  }
}
