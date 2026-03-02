resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[1]s"
  type             = "self_hosted"
  domain           = "%[1]s.%[3]s"
  session_duration = "24h"

  oauth_configuration = {
    enabled = true

    dynamic_client_registration = {
      enabled               = true
      allow_any_on_localhost = false
      allow_any_on_loopback = true
      allowed_uris          = ["https://example.com/callback", "https://app.example.com/oauth/*"]
    }

    grant = {
      access_token_lifetime = "30m"
      session_duration      = "12h"
    }
  }
}
