resource "cloudflare_api_shield_operation" "%s" {
  zone_id  = "%s"
  method   = "GET"
  host     = "api.example.com"
  endpoint = "/api/v1/users/{var1}/posts/{var2}"
}
