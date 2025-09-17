resource "cloudflare_bot_management" "%[1]s" {
	zone_id = "%[2]s"

	# Configuration that might be SBFM plan where suppress_session_score isn't returned by API
	sbfm_definitely_automated = "block"
	sbfm_verified_bots = "allow"
	sbfm_static_resource_protection = true
	optimize_wordpress = false

	# This field might not be returned by API for SBFM plans
	# but Terraform expects it with default false
	# This should trigger the drift issue
}