resource "cloudflare_bot_management" "%[1]s" {
	zone_id = "%[2]s"

	sbfm_definitely_automated = "block"
	sbfm_likely_automated = "managed_challenge"
	sbfm_verified_bots = "allow"
	sbfm_static_resource_protection = true
}