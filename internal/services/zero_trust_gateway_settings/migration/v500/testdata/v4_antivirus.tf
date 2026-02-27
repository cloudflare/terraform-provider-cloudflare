resource "cloudflare_teams_account" "%[1]s" {
  account_id                 = "%[2]s"
  activity_log_enabled       = false
  tls_decrypt_enabled        = false
  protocol_detection_enabled = false

  antivirus {
    enabled_download_phase = true
    enabled_upload_phase   = false
    fail_closed            = true

    notification_settings {
      enabled     = true
      message     = "File scanning in progress"
      support_url = "https://support.example.com/"
    }
  }
}
