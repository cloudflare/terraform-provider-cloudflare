resource "cloudflare_url_normalization_settings" "%s" {
  zone_id = "%s"
  type    = "rfc3986"
  scope   = "both"
}
