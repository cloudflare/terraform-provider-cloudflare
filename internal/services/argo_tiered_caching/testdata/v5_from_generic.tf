resource "cloudflare_argo_tiered_caching" "%[1]s" {
  zone_id = "%[2]s"
  value   = "on"
}

moved {
  from = cloudflare_tiered_cache.%[1]s
  to   = cloudflare_argo_tiered_caching.%[1]s
}