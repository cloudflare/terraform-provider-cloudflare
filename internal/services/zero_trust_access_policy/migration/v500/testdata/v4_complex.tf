resource "cloudflare_access_policy" "%[1]s" {
  account_id                     = "%[2]s"
  name                           = "%[1]s"
  decision                       = "allow"
  session_duration               = "24h"
  approval_required              = true
  purpose_justification_required = true
  purpose_justification_prompt   = "Why do you need access?"

  approval_group {
    approvals_needed = 2
    email_addresses  = ["admin1@example.com", "admin2@example.com"]
  }

  include {
    email                   = ["user1@example.com", "user2@example.com"]
    email_domain            = ["example.com", "test.com"]
    ip                      = ["192.168.1.0/24", "10.0.0.0/8"]
    everyone                = true
    any_valid_service_token = true
  }

  exclude {
    email = ["blocked@example.com"]
    geo   = ["CN", "RU"]
  }

  require {
    certificate = true
  }
}
