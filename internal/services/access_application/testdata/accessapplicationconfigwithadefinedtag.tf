
resource "cloudflare_access_tag" "%[1]s" {
  account_id = "%[4]s"
  name = "%[1]s"
}
resource "cloudflare_access_application" "%[1]s" {
  zone_id                   = "%[2]s"
  name                      = "%[1]s"
  domain                    = "%[1]s.%[3]s"
  type                      = "self_hosted"
  session_duration          = "24h"
  tags                      = [cloudflare_access_tag.%[1]s.id]
}
