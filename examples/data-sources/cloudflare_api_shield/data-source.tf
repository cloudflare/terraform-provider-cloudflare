data "cloudflare_api_shield" "example_api_shield" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  properties = ["auth_id_characteristics"]
}
