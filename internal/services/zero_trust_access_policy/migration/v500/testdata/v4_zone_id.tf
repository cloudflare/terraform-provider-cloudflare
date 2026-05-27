resource "cloudflare_access_application" "%[1]s" {
  zone_id          = "%[2]s"
  name             = "%[1]s"
  domain           = "%[1]s.%[3]s"
  session_duration = "24h"
}

resource "cloudflare_access_policy" "%[1]s" {
  zone_id          = "%[2]s"
  application_id   = cloudflare_access_application.%[1]s.id
  name             = "%[1]s"
  decision         = "allow"
  precedence       = 1
  session_duration = "24h"
  include {
    email = ["test@example.com"]
  }
}
