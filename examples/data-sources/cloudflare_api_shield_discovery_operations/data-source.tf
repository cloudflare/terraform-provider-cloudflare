data "cloudflare_api_shield_discovery_operations" "example_api_shield_discovery_operations" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  diff = true
  direction = "desc"
  endpoint = "/api/v1"
  host = ["api.cloudflare.com"]
  method = ["GET"]
  order = "method"
  origin = "ML"
  state = "review"
}
