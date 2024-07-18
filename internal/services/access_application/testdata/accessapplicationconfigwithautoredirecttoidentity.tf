
resource "cloudflare_access_identity_provider" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[1]s"
  type    = "onetimepin"
}

resource "cloudflare_access_application" "%[1]s" {
  zone_id                   = "%[2]s"
  name                      = "%[1]s"
  domain                    = "%[1]s.%[3]s"
  type                      = "self_hosted"
  session_duration          = "24h"
  auto_redirect_to_identity = true
  allowed_idps              = [cloudflare_access_identity_provider.%[1]s.id]

  depends_on = ["cloudflare_access_identity_provider.%[1]s"]
}
