# Enable the Cache Reserve support for a given zone.
resource "cloudflare_zone_cache_variants" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  enabled = true
}
