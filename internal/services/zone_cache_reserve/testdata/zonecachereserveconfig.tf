
		resource "cloudflare_zone_cache_reserve" "%[2]s" {
			zone_id = "%[1]s"
			enabled = "%[3]t"
		}