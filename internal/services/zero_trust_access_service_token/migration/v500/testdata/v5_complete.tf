resource "cloudflare_zero_trust_access_service_token" "%s" {
  account_id = "%s"
  name       = "test-%s"
  duration   = "43800h"
}
