resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id                = "%[2]s"
  name                      = "%[1]s"
  type                      = "saas"
  session_duration          = "24h"
  auto_redirect_to_identity = false
}
