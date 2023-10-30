resource "cloudflare_api_shield_schema_validation_settings" "example" {
  zone_id                               = "0da42c8d2132a9ddaf714f9e7c920711"
  validation_default_mitigation_action  = "log"
  validation_override_mitigation_action = "none"
}
