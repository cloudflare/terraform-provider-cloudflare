
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  %[3]s_id                  = "%[4]s"
  name                      = "%[1]s"
  type                      = "self_hosted"
  session_duration          = "24h"
  auto_redirect_to_identity = false
  self_hosted_domains       = [
    "d1.%[1]s.%[2]s",
    "d2.%[1]s.%[2]s"
  ]
}
