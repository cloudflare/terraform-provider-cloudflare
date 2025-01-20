resource "cloudflare_zero_trust_gateway_settings" "%[1]s" {
  account_id = "%[2]s"
  settings = {
    protocol_detection = {
        enabled = true
    }
    tls_decrypt = {
      enabled = true
    }
    activity_log = {
      enabled = true
    }
    browser_isolation = {
      url_browser_isolation_enabled = true
      non_identity_enabled = false
    }
    block_page = {
      name = "%[1]s"
      enabled = true
      footer_text = "hello"
      header_text = "hello"
      logo_path = "https://example.com"
      background_color = "#000000"
      mailto_subject = "hello"
      mailto_address = "test@cloudflare.com"
    }
    body_scanning = {
      inspection_mode = "deep"
    }
    fips = {
      tls = true
    }
    antivirus = {
      enabled_download_phase = true
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
    custom_certificate = {
      enabled = false
    }
  }
}
