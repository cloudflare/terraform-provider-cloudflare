resource "cloudflare_zero_trust_gateway_settings" "%[1]s" {
  account_id = "%[2]s"
  settings = {
    protocol_detection = {
      enabled = false
    }
    tls_decrypt = {
      enabled = false
    }
    activity_log = {
      enabled = false
    }
    browser_isolation = {
      url_browser_isolation_enabled = false
      non_identity_enabled = true
    }
    body_scanning = {
      inspection_mode = "shallow"
    }
    fips = {
      tls = false
    }
    antivirus = {
      enabled_download_phase = false
      enabled_upload_phase = true
      fail_closed = false
      notification_settings = {
        enabled = false
        msg = "updated msg"
        support_url = "https://updated.com/"
      }
    }
    extended_email_matching = {
      enabled = false
    }
  }
}