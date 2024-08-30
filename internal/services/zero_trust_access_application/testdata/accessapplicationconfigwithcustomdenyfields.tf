
resource "cloudflare_zero_trust_access_application" "%[1]s" {
  zone_id                   = "%[2]s"
  name                      = "%[1]s"
  domain                    = "%[1]s.%[3]s"
  type                      = "self_hosted"
  session_duration          = "24h"
  custom_deny_message       = "denied!"
  custom_deny_url           = "https://www.cloudflare.com"
	custom_non_identity_deny_url = "https://www.blocked.com"
}
