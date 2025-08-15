resource "cloudflare_bot_management" "%[1]s" {
	zone_id = "%[2]s"
	
	# Don't set auto_update_model explicitly
	# API returns null but schema expects default true
	enable_js = true
}