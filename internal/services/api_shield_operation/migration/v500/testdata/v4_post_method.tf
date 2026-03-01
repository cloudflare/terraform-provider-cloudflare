resource "cloudflare_api_shield_operation" "%s" {
  zone_id  = "%s"
  method   = "POST"
  host     = "api.example.com"
  endpoint = "/api/v1/users"
}
