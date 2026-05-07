resource "cloudflare_zero_trust_organization" "%[1]s" {
  account_id  = "%[3]s"
  name        = "%[4]s"
  auth_domain = "%[1]s-%[4]s"

  mfa_config = {
    allowed_authenticators = ["totp", "security_key"]
    session_duration       = "24h"
  }
}

resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id       = "%[3]s"
  name             = "%[1]s"
  type             = "self_hosted"
  domain           = "%[1]s.%[2]s"
  session_duration = "24h"

  mfa_config = {
    allowed_authenticators = ["totp", "security_key"]
    mfa_disabled           = false
    session_duration       = "12h"
  }

  depends_on = [cloudflare_zero_trust_organization.%[1]s]
}
