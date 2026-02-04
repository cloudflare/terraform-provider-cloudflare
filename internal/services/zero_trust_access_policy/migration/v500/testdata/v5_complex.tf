resource "cloudflare_zero_trust_access_policy" "%[1]s" {
  account_id                     = "%[2]s"
  name                           = "%[1]s"
  decision                       = "allow"
  session_duration               = "24h"
  approval_required              = true
  purpose_justification_required = true
  purpose_justification_prompt   = "Why do you need access?"

  approval_groups = [
    {
      approvals_needed = 2
      email_addresses  = ["admin1@example.com", "admin2@example.com"]
    }
  ]

  include = [
    { email = { email = "user1@example.com" } },
    { email = { email = "user2@example.com" } },
    { everyone = {} }
  ]

  exclude = [
    { email = { email = "blocked@example.com" } }
  ]
}
