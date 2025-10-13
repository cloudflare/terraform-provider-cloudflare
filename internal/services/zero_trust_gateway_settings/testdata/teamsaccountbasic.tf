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
      enabled = true
      background_color = "#000000"
      footer_text = "footer"
      header_text = "header"
      logo_path = "https://example.com"
      name = "%[1]s"
      include_context = true
      mailto_address = "test@cloudflare.com"
      mailto_subject = "hello"
      target_uri = ""
      suppress_footer = false
      mode = "customized_block_page"
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
    host_selector = {
      enabled = false
    }
    inspection = {
      mode = "static"
    }
  }
}
