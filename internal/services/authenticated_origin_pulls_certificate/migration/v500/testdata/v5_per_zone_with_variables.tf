# v5 per-zone certificate using file() references
# type field removed in v5
resource "cloudflare_authenticated_origin_pulls_certificate" {
  zone_id     = "%s"
  certificate = file("${path.module}/cert.pem")
  private_key = file("${path.module}/key.pem")
}
