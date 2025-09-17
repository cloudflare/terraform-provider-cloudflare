resource "cloudflare_argo_tiered_caching" "%[2]s" {
	zone_id = "%[1]s"
  value   = "on"
}
