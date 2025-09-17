resource "cloudflare_bot_management" "%[1]s" {
	zone_id = "%[2]s"

	auto_update_model = %[3]t
}