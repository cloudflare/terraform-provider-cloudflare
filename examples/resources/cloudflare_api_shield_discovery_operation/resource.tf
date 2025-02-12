resource "cloudflare_api_shield_discovery_operation" "example_api_shield_discovery_operation" {
  zone_id = "zone_id"
  operation_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  state = "review"
}
