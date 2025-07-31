resource "cloudflare_argo_tiered_caching" "%[2]s" {
	zone_id = "%[1]s"
	value   = "on"
}

data "cloudflare_argo_tiered_caching" "%[2]s" {
	zone_id = cloudflare_argo_tiered_caching.%[2]s.zone_id
	
	depends_on = [cloudflare_argo_tiered_caching.%[2]s]
}