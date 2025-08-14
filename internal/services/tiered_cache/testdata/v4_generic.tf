resource "cloudflare_tiered_cache" "%[1]s" {
  zone_id    = "%[2]s"
  cache_type = "generic"
}