data "cloudflare_api_shield_operation" "example_api_shield_operation" {
  endpoint = "/api/v1"
  host = ["api.cloudflare.com"]
  method = ["GET"]
  origin = []
  state = "review"
}
