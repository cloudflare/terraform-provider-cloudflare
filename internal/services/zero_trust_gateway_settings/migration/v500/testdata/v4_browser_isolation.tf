resource "cloudflare_teams_account" "%[1]s" {
  account_id                             = "%[2]s"
  activity_log_enabled                   = false
  tls_decrypt_enabled                    = false
  protocol_detection_enabled             = false
  url_browser_isolation_enabled          = true
  non_identity_browser_isolation_enabled = false
}
