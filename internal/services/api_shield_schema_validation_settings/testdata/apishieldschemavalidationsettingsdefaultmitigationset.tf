
	resource "cloudflare_api_shield_schema_validation_settings" "%[1]s" {
		zone_id = "%[2]s"
		validation_default_mitigation_action = "log"
	}
