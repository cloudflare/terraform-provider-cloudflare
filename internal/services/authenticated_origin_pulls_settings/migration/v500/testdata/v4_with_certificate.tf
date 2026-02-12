resource "cloudflare_authenticated_origin_pulls" "%[1]s" {
  zone_id                                = "%[2]s"
  authenticated_origin_pulls_certificate = "cert-123"
  enabled                                = true
}
