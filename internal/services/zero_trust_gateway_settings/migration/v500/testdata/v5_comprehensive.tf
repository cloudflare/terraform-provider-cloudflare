resource "cloudflare_zero_trust_gateway_settings" "%[1]s" {
  account_id = "%[2]s"
  settings = {
    activity_log = {
      enabled = true
    }
    tls_decrypt = {
      enabled = true
    }
    certificate = {
      id = "%[3]s"
    }
    protocol_detection = {
      enabled = false
    }
    browser_isolation = {
      url_browser_isolation_enabled = true
      non_identity_enabled          = false
    }
    fips = {
      tls = true
    }
    body_scanning = {
      inspection_mode = "deep"
    }
    antivirus = {
      enabled_download_phase = true
      enabled_upload_phase   = false
      fail_closed            = true
      notification_settings = {
        enabled     = true
        msg         = "Scanning"
        support_url = "https://support.example.com/"
      }
    }
  }
}
