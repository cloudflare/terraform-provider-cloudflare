data "cloudflare_api_shield" "example_api_shield" {
  zone_id = "zone_id"
  properties = ["auth_id_characteristics"]
}
