resource "cloudflare_zero_trust_gateway_settings" "%[1]s" {
  account_id = "%[2]s"
  settings = {
    activity_log = {
      enabled = false
    }
    tls_decrypt = {
      enabled = false
    }
    protocol_detection = {
      enabled = false
    }
    antivirus = {
      enabled_download_phase = true
      enabled_upload_phase   = false
      fail_closed            = true
      notification_settings = {
        enabled     = true
        msg         = "File scanning in progress"
        support_url = "https://support.example.com/"
      }
    }
  }
}
