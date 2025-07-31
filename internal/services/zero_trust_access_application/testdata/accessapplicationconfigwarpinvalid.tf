resource "cloudflare_zero_trust_access_application" "%[1]s" {
  name                        = "Warp Login App"
  account_id                  = "c91137350c00a28806157ec3918faff2"
  type                        = "warp"
  allow_authenticate_via_warp = false
  session_duration            = "3h"
  policies = []
}