resource "cloudflare_teams_account" "%[1]s" {
  account_id                 = "%[2]s"
  activity_log_enabled       = true
  tls_decrypt_enabled        = false
  protocol_detection_enabled = false
}
