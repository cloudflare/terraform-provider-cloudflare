
    resource "cloudflare_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_access_policy" "%[1]s" {
      application_id = "${cloudflare_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "allow"
      precedence     = "1"

      include =[ {
        email = ["a@example.com", "b@example.com"]
      }]

      purpose_justification_required = "true"
      purpose_justification_prompt = "Why should we let you in?"
    }
  