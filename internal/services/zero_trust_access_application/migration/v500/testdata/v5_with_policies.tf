moved {
  from = cloudflare_access_policy.%[1]s
  to   = cloudflare_zero_trust_access_policy.%[1]s
}

moved {
  from = cloudflare_access_application.%[1]s
  to   = cloudflare_zero_trust_access_application.%[1]s
}

resource "cloudflare_zero_trust_access_policy" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s-policy"
  decision   = "allow"
  include = [{
    everyone = true
  }]
}

resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  domain     = "%[1]s.%[3]s"
  type       = "self_hosted"
  session_duration = "24h"

  policies = [{
    id = cloudflare_zero_trust_access_policy.%[1]s.id
  }]
}
