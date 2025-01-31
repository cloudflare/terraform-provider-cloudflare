resource "cloudflare_api_shield" "%[1]s" {
  zone_id = "%[2]s"
  auth_id_characteristics = []
}