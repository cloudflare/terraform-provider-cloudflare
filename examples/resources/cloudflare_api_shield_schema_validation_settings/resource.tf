resource "cloudflare_api_shield_schema_validation_settings" "example_api_shield_schema_validation_settings" {
  zone_id = "zone_id"
  validation_default_mitigation_action = "none"
  validation_override_mitigation_action = "none"
}
