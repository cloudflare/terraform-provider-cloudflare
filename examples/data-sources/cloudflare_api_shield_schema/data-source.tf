data "cloudflare_api_shield_schema" "example_api_shield_schema" {
  zone_id = "zone_id"
  schema_id = "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"
  omit_source = true
}
