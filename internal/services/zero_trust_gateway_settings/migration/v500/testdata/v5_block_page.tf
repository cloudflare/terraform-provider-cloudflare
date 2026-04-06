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
    block_page = {
      enabled          = true
      name             = "%[1]s"
      footer_text      = "Contact IT"
      header_text      = "Blocked"
      logo_path        = "https://example.com/logo.png"
      background_color = "#FF0000"
      mailto_address   = "security@example.com"
      mailto_subject   = "Access Request"
    }
  }
}
