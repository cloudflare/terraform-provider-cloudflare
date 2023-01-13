data "cloudflare_access_identity_provider" "example" {
  name       = "Google SSO"
  account_id = "example-account-id"
}

resource "cloudflare_access_application" "example" {
  zone_id                   = "example.com"
  name                      = "name"
  domain                    = "name.example.com"
  type                      = "self_hosted"
  session_duration          = "24h"
  allowed_idps              = [data.cloudflare_access_identity_provider.example.id]
  auto_redirect_to_identity = true
}
