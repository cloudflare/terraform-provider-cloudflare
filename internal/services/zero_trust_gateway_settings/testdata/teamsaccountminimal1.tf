resource "cloudflare_zero_trust_gateway_settings" "%[1]s" {
  account_id = "%[2]s"
  settings = {
    protocol_detection = {
        enabled = false
    }
    tls_decrypt = {
      enabled = true
    }
    activity_log = {
      enabled = true
    }

    antivirus = {
      enabled_download_phase = false
      enabled_upload_phase = false
      fail_closed = true
      notification_settings = {
          enabled = true
          msg = "msg"
          support_url = "https://hello.com/"
      }
    }
    extended_email_matching = {
      enabled = true
    }
  }
}
