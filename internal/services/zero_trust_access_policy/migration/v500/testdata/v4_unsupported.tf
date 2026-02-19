resource "cloudflare_access_policy" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[1]s"
  decision         = "allow"
  session_duration = "24h"

  include {
    everyone = true
  }

  exclude {
    email = ["test@blocked.com"]
  }
}
