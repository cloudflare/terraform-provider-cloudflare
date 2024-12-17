
	resource "cloudflare_bot_management" "%[1]s" {
		zone_id = "%[2]s"

		enable_js = "%[3]t"

		sbfm_definitely_automated = "%[4]s"
		sbfm_likely_automated = "%[5]s"
		sbfm_verified_bots = "%[6]s"
		sbfm_static_resource_protection = "%[7]t"
		optimize_wordpress = "%[8]t"
	}
