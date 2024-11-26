# Per-Zone Authenticated Origin Pulls certificate
resource "cloudflare_authenticated_origin_pulls_certificate" "my_per_zone_aop_cert" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  certificate = "-----INSERT CERTIFICATE-----"
  private_key = "-----INSERT PRIVATE KEY-----"
  type        = "per-zone"
}

# Per-Hostname Authenticated Origin Pulls certificate
resource "cloudflare_authenticated_origin_pulls_certificate" "my_per_hostname_aop_cert" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  certificate = "-----INSERT CERTIFICATE-----"
  private_key = "-----INSERT PRIVATE KEY-----"
  type        = "per-hostname"
}
