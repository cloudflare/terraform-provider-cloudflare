resource "cloudflare_argo" "%s" {
  zone_id         = "%s"
  smart_routing   = "on"
  tiered_caching  = "on"
}
