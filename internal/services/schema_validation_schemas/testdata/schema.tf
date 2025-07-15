resource "cloudflare_api_shield_operation" "getAllProducts" {
  zone_id = "%[2]s"
  endpoint = "/products"
  host = "api.example.com"
  method = "GET"
}

resource "cloudflare_schema_validation_schemas" "%[1]s" {
  zone_id = "%[2]s"
  kind = "openapi_v3"
  name = "%[1]s.yaml"
  source = file("testdata/test_schema.yaml")
  validation_enabled = %[3]s
}
