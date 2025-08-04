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
      mailto_address = "test@cloudflare.com"
      mailto_subject = "hello"
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
    }
    extended_email_matching = {
      enabled = true
    }
  }
}