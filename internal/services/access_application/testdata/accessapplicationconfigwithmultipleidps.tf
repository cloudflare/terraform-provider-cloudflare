
resource "cloudflare_access_identity_provider" "%[5]s" {
  account_id = "%[4]s"
  name = "%[5]s"
  type = "onetimepin"
}

resource "cloudflare_access_identity_provider" "%[6]s" {
  account_id = "%[4]s"
  name = "%[6]s"
  type = "github"
  config = {
  client_id = "test"
    client_secret = "secret"
}
}

resource "cloudflare_access_application" "%[1]s" {
  zone_id                   = "%[2]s"
  name                      = "%[1]s"
  domain                    = "%[1]s.%[3]s"
  type                      = "self_hosted"
  session_duration          = "24h"
  allowed_idps              = [
    cloudflare_access_identity_provider.%[5]s.id,
    cloudflare_access_identity_provider.%[6]s.id,
  ]
}
