
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  zone_id                    = "%[2]s"
  name                       = "%[1]s"
  domain                     = "%[1]s.%[3]s"
  type                       = "self_hosted"
  session_duration           = "24h"
  http_only_cookie_attribute = false
}
