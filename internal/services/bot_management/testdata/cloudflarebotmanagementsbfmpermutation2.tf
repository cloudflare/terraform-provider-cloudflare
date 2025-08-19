resource "cloudflare_bot_management" "%[1]s" {
	zone_id = "%[2]s"

	sbfm_definitely_automated = "managed_challenge"
	sbfm_likely_automated = "allow"
	sbfm_verified_bots = "allow"
	sbfm_static_resource_protection = false
}