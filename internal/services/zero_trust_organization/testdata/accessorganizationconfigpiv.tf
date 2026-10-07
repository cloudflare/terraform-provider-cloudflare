resource "cloudflare_zero_trust_organization" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[4]s"
  auth_domain      = "%[1]s-%[3]s"
  session_duration = "12h"

  mfa_piv_key_requirements = {
    pin_policy          = "once"
    require_fips_device = true
    ssh_key_size        = [256]
    ssh_key_type        = ["ecdsa"]
    touch_policy        = "always"
  }
}

data "cloudflare_zero_trust_organization" "%[1]s" {
  account_id = cloudflare_zero_trust_organization.%[1]s.account_id
}
