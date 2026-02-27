resource "cloudflare_teams_account" "%[1]s" {
  account_id                 = "%[2]s"
  activity_log_enabled       = false
  tls_decrypt_enabled        = false
  protocol_detection_enabled = false

  fips {
    tls = true
  }

  body_scanning {
    inspection_mode = "deep"
  }
}
