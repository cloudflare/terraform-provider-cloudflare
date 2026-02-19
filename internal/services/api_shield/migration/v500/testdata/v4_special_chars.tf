resource "cloudflare_api_shield" "%s" {
  zone_id = "%s"

  auth_id_characteristics {
    type = "header"
    name = "X-API-Key"
  }

  auth_id_characteristics {
    type = "header"
    name = "X_Custom_Header"
  }

  auth_id_characteristics {
    type = "cookie"
    name = "SessionID"
  }

  auth_id_characteristics {
    type = "cookie"
    name = "user-session-token"
  }
}
