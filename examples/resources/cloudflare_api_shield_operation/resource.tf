resource "cloudflare_api_shield_operation" "example" {
  zone_id  = "0da42c8d2132a9ddaf714f9e7c920711"
  method   = "GET"
  host     = "api.example.com"
  endpoint = "/path"
}
