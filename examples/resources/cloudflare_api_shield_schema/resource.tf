resource "cloudflare_api_shield_schema" "example_api_shield_schema" {
  zone_id = "zone_id"
  file = "file.txt"
  kind = "openapi_v3"
  name = "petstore schema"
  validation_enabled = "true"
}
