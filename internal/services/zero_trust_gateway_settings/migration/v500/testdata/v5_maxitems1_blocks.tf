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
    fips = {
      tls = true
    }
    body_scanning = {
      inspection_mode = "deep"
    }
  }
}
