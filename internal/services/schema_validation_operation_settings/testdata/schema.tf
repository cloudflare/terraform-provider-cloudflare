resource "cloudflare_api_shield_operation" "getAllProductsOne" {
  zone_id = "%[2]s"
  endpoint = "/products_one"
  host = "api.example.com"
  method = "GET"
}

resource "cloudflare_schema_validation_schemas" "%[1]s" {
  zone_id = "%[2]s"
  kind = "openapi_v3"
  name = "test_schema.yaml"
  source = file("testdata/test_schema.yaml")
  validation_enabled = true
}

resource "cloudflare_schema_validation_operation_settings" "%[1]s" {
  zone_id = "%[2]s"
  operation_id = cloudflare_api_shield_operation.getAllProductsOne.operation_id
  mitigation_action = "%[3]s"
}
