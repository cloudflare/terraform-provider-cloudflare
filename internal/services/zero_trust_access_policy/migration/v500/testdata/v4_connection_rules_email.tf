resource "cloudflare_access_policy" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[1]s"
  decision         = "allow"
  session_duration = "24h"

  include {
    everyone = true
  }

  connection_rules {
    ssh {
      usernames         = ["root", "admin"]
      allow_email_alias = true
    }
  }
}
