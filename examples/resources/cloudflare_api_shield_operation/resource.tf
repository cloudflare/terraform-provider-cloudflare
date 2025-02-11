resource "cloudflare_api_shield_operation" "example_api_shield_operation" {
  zone_id = "zone_id"
  endpoint = "/api/v1/users/{var1}"
  host = "www.example.com"
  method = "GET"
}
