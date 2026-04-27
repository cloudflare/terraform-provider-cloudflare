resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "bookmark"

  mfa_config = {
    allowed_authenticators = ["totp"]
    mfa_disabled           = false
    session_duration       = "1h"
  }
}
