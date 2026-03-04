# v4 per-zone certificate using file() references
# Tests that migration handles expressions correctly
resource "cloudflare_authenticated_origin_pulls_certificate" {
  zone_id     = "%s"
  certificate = file("${path.module}/cert.pem")
  private_key = file("${path.module}/key.pem")
  type        = "per-zone"
}
