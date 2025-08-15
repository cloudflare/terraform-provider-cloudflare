resource "cloudflare_bot_management" "%[1]s" {
	zone_id = "%[2]s"

	# Minimal configuration - API returns null for fight_mode and optimize_wordpress
	# but schema expects defaults, causing drift
}