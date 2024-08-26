
    resource "cloudflare_zero_trust_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_zero_trust_access_policy" "%[1]s" {
      application_id = cloudflare_zero_trust_access_application.%[1]s.id
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "allow"
      precedence     = "1"

      include =[ {
		external_evaluation = {
			evaluate_url = "https://example.com"
			keys_url = "https://example.com/keys"
		  }]
      }]
    }
