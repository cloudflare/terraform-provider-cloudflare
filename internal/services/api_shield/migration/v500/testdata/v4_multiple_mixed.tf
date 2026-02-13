resource "cloudflare_api_shield" "%s" {
  zone_id = "%s"

  auth_id_characteristics {
    type = "header"
    name = "authorization"
  }

  auth_id_characteristics {
    type = "cookie"
    name = "session_id"
  }

  auth_id_characteristics {
    type = "header"
    name = "x-api-key"
  }
}
