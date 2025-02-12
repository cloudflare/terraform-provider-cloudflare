data "cloudflare_api_shield_operation" "example_api_shield_operation" {
  zone_id = "zone_id"
  operation_id = "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"
  feature = ["thresholds"]
}
