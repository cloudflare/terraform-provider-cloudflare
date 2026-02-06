resource "cloudflare_api_shield" "%s" {
  zone_id = "%s"

  auth_id_characteristics {
    type = "header"
    name = "authorization"
  }
}
