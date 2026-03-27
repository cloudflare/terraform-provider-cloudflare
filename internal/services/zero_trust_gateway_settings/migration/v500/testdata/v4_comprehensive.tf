resource "cloudflare_teams_account" "%[1]s" {
  account_id                             = "%[2]s"
  activity_log_enabled                   = true
  tls_decrypt_enabled                    = false
  protocol_detection_enabled             = false
  url_browser_isolation_enabled          = true
  non_identity_browser_isolation_enabled = false

  fips {
    tls = true
  }

  body_scanning {
    inspection_mode = "deep"
  }

  antivirus {
    enabled_download_phase = true
    enabled_upload_phase   = false
    fail_closed            = true

    notification_settings {
      enabled     = true
      message     = "Scanning"
      support_url = "https://support.example.com/"
    }
  }
}
