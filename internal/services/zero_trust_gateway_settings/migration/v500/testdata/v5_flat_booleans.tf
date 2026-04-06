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
  }
}
