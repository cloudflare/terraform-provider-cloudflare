resource "cloudflare_bot_management" "%[1]s" {
	zone_id = "%[2]s"
	
	ai_bots_protection = "block"
	crawler_protection = "enabled"
	enable_js = true

	lifecycle {
		ignore_changes = [
			auto_update_model,
			optimize_wordpress,
			sbfm_definitely_automated,
			sbfm_likely_automated,
			sbfm_static_resource_protection,
			sbfm_verified_bots,
			stale_zone_configuration,
			suppress_session_score,
			using_latest_model,
		]
	}
}