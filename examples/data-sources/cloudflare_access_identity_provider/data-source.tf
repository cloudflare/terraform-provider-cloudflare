data "cloudflare_access_identity_provider" "example" {
  name       = "Google SSO"
  account_id = "f037e56e89293a057740de681ac9abbe"
}

resource "cloudflare_access_application" "example" {
  zone_id                   = "0da42c8d2132a9ddaf714f9e7c920711"
  name                      = "name"
  domain                    = "name.example.com"
  type                      = "self_hosted"
  session_duration          = "24h"
  allowed_idps              = [data.cloudflare_access_identity_provider.example.id]
  auto_redirect_to_identity = true
}
