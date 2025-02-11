data "cloudflare_api_shield_discovery_operations" "example_api_shield_discovery_operations" {
  zone_id = "zone_id"
  diff = true
  direction = "asc"
  endpoint = "/api/v1"
  host = ["api.cloudflare.com"]
  method = ["GET"]
  order = "host"
  origin = "ML"
  state = "review"
}
