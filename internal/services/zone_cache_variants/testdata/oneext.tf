
		resource "cloudflare_zone_cache_variants" "%[2]s" {
			zone_id = "%[1]s"
			avif = ["image/avif", "image/webp"]
		}