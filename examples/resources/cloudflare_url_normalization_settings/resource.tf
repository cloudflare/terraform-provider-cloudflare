resource "cloudflare_url_normalization_settings" "example_url_normalization_settings" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  scope = "incoming"
  type = "cloudflare"
}
