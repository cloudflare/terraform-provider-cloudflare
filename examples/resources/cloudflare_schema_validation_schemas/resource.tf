resource "cloudflare_schema_validation_schemas" "example_schema_validation_schemas" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  kind = "openapi_v3"
  name = "petstore schema"
  source = "<schema file contents>"
  validation_enabled = true
}
