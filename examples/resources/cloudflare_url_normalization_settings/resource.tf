resource "cloudflare_url_normalization_settings" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  type    = "cloudflare"
  scope   = "incoming"
}