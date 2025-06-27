resource "cloudflare_zero_trust_access_policy" "%[1]s_reusable_p1" {
  account_id = "%[3]s"
  name       = "%[1]s"
  decision   = "non_identity"
  include = [
    {
      ip = { ip = "127.0.0.1/32" }
    }
  ]
}

resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[3]s"
  name       = "%[1]s"
  domain     = "%[1]s.%[2]s"
  type       = "self_hosted"
  policies = [
    {
      name     = "%[1]s.legacy_p1"
      decision = "allow"
      include = [
        {
          email = { email = "a@example.com" }
        }
      ]
      precedence = 2
    },
    {
      id         = cloudflare_zero_trust_access_policy.%[1]s_reusable_p1.id
      precedence = 4
    },
    {
      name     = "%[1]s.legacy_p2"
      decision = "non_identity"
      include = [
        {
          ip = { ip = "127.0.0.1/32" }
        }
      ]
    }
  ]
}
