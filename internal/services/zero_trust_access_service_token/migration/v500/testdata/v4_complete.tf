resource "cloudflare_zero_trust_access_service_token" "%s" {
  account_id           = "%s"
  name                 = "test-%s"
  duration             = "43800h"
  min_days_for_renewal = 60
}
