resource "cloudflare_custom_hostname" "%s" {
  zone_id  = "%s"
  hostname = "%s.%s"

  ssl {
    method   = "txt"
    type     = "dv"

    settings {
      http2           = "on"
      tls13           = "on"
      min_tls_version = "1.2"
      early_hints     = "off"
      ciphers         = ["ECDHE-RSA-AES128-GCM-SHA256"]
    }
  }

  custom_origin_server = "origin-%s.%s"
  custom_origin_sni    = "origin-%s.%s"
  custom_metadata = {
    environment = "migration"
    owner       = "terraform"
  }

  wait_for_ssl_pending_validation = false

  lifecycle {
    ignore_changes = [
      ownership_verification,
      ownership_verification_http,
      status,
    ]
  }
}
