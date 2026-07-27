resource "cloudflare_zero_trust_organization" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[3]s"
  auth_domain      = "%[1]s-%[3]s"
  session_duration = "12h"

  mfa_config = {
    allowed_authenticators      = ["totp", "security_key"]
    session_duration            = "24h"
    amr_matching_session_duration = "12h"
  }
}
