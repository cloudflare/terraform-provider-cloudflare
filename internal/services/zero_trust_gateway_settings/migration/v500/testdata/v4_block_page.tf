resource "cloudflare_teams_account" "%[1]s" {
  account_id                 = "%[2]s"
  activity_log_enabled       = false
  tls_decrypt_enabled        = false
  protocol_detection_enabled = false

  block_page {
    enabled          = true
    name             = "%[1]s"
    footer_text      = "Contact IT"
    header_text      = "Blocked"
    logo_path        = "https://example.com/logo.png"
    background_color = "#FF0000"
    mailto_address   = "security@example.com"
    mailto_subject   = "Access Request"
  }
}
