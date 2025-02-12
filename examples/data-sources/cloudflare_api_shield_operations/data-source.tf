data "cloudflare_api_shield_operations" "example_api_shield_operations" {
  zone_id = "zone_id"
  direction = "asc"
  endpoint = "/api/v1"
  feature = ["thresholds"]
  host = ["api.cloudflare.com"]
  method = ["GET"]
  order = "method"
}
