resource "cloudflare_bot_management" "%[1]s" {
	zone_id = "%[2]s"

	# Only set fields that are available with current entitlements
	# Don't set optimize_wordpress or fight_mode - API returns null
	# but schema has default values
	enable_js = true
}