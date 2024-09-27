resource "cloudflare_api_shield_operation" "example" {
  zone_id  = "0da42c8d2132a9ddaf714f9e7c920711"
  method   = "GET"
  host     = "api.example.com"
  endpoint = "/path"
}

resource "cloudflare_api_shield_operation_schema_validation_settings" "example" {
  zone_id           = "0da42c8d2132a9ddaf714f9e7c920711"
  operation_id      = cloudflare_api_shield_operation.example.id
  mitigation_action = "block"
}
