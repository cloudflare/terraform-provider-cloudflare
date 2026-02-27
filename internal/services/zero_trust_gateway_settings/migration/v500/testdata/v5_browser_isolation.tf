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
    browser_isolation = {
      url_browser_isolation_enabled = true
      non_identity_enabled          = false
    }
  }
}
