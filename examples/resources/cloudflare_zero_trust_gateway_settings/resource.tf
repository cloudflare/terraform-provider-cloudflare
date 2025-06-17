resource "cloudflare_zero_trust_gateway_settings" "example_zero_trust_gateway_settings" {
  account_id = "699d98642c564d2e855e9661899b7252"
  settings = {
    activity_log = {
      enabled = true
    }
    antivirus = {
      enabled_download_phase = false
      enabled_upload_phase = false
      fail_closed = false
      notification_settings = {
        enabled = true
        include_context = true
        msg = "msg"
        support_url = "support_url"
      }
    }
    app_control_settings = {
      enabled = false
    }
    block_page = {
      background_color = "background_color"
      enabled = true
      footer_text = "--footer--"
      header_text = "--header--"
      include_context = true
      logo_path = "https://logos.com/a.png"
      mailto_address = "admin@example.com"
      mailto_subject = "Blocked User Inquiry"
      mode = "customized_block_page"
      name = "Cloudflare"
      suppress_footer = false
      target_uri = "https://example.com"
    }
    body_scanning = {
      inspection_mode = "deep"
    }
    browser_isolation = {
      non_identity_enabled = true
      url_browser_isolation_enabled = true
    }
    certificate = {
      id = "d1b364c5-1311-466e-a194-f0e943e0799f"
    }
    custom_certificate = {
      enabled = true
      id = "d1b364c5-1311-466e-a194-f0e943e0799f"
    }
    extended_email_matching = {
      enabled = true
    }
    fips = {
      tls = true
    }
    host_selector = {
      enabled = false
    }
    protocol_detection = {
      enabled = true
    }
    sandbox = {
      enabled = true
      fallback_action = "allow"
    }
    tls_decrypt = {
      enabled = true
    }
  }
}
