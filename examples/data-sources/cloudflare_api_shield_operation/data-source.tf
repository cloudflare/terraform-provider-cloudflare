data "cloudflare_api_shield_operation" "example_api_shield_operation" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  diff = true
  direction = "asc"
  endpoint = "/api/v1"
  host = ["api.cloudflare.com"]
  method = ["GET"]
  order = "host"
  origin = "ML"
  page = 1
  per_page = 5
  state = "review"
}
