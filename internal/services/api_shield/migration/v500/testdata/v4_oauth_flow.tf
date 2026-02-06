resource "cloudflare_api_shield" "%s" {
  zone_id = "%s"

  auth_id_characteristics {
    type = "header"
    name = "Authorization"
  }

  auth_id_characteristics {
    type = "header"
    name = "X-OAuth-Token"
  }

  auth_id_characteristics {
    type = "cookie"
    name = "oauth_state"
  }

  auth_id_characteristics {
    type = "header"
    name = "X-Request-ID"
  }
}
