
resource "cloudflare_custom_hostname" "%[2]s" {
  zone_id = "%[1]s"
  hostname = "%[2]s.%[3]s"
  ssl = {
    method = "http"
    type = "dv"
    settings = {
      http2 = "off"
      min_tls_version = "1.1"
      ciphers = [
        "ECDHE-RSA-AES128-GCM-SHA256",
        "AES128-SHA"
      ]
      early_hints = "off"
    }
    certificate_authority = "google"
  }
  
  lifecycle {
    ignore_changes = [
      created_at,
      ownership_verification,
      ownership_verification_http,
      ssl.wildcard,
      status,
      verification_errors
    ]
  }
}
