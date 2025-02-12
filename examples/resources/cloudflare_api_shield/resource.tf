resource "cloudflare_api_shield" "example_api_shield" {
  zone_id = "zone_id"
  auth_id_characteristics = [{
    name = "authorization"
    type = "header"
  }]
}
