resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[1]s"
  type             = "self_hosted"
  domain           = "%[1]s.%[3]s"
  session_duration = "24h"

  oauth_configuration = {
    enabled = true
  }
}
