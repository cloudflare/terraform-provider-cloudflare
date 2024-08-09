
    resource "cloudflare_zero_trust_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_zero_trust_access_policy" "%[1]s" {
      application_id = "${cloudflare_zero_trust_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "allow"
      precedence     = "1"

      purpose_justification_required = "true"
      purpose_justification_prompt = "Why should we let you in?"
      approval_required = "true"

      include =[ {
        email = ["a@example.com", "b@example.com"]
      }]

      approval_group =[ {
        email_addresses = ["test1@example.com", "test2@example.com", "test3@example.com"]
        approvals_needed = "2"
      },
    {
    email_addresses = ["test4@example.com"]
        approvals_needed = "1"
    }]

    }
  