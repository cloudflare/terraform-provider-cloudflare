resource "cloudflare_zero_trust_access_policy" "%[1]s_p1" {
  account_id 	= "%[3]s"
  name          = "%[1]s"
  decision	= "allow"
  include = [{
    email = { email = "a@example.com" }
  }]
}

resource "cloudflare_zero_trust_access_policy" "%[1]s_p2" {
  account_id	= "%[3]s"
  name          = "%[1]s"
  decision	= "non_identity"
  include = [{
    ip = { ip = "127.0.0.1/32" }
  }]
}

resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id   = "%[3]s"
  name         = "%[1]s"
  domain       = "%[1]s.%[2]s"
  type         = "self_hosted"
  policies     = [
    { 
      id = cloudflare_zero_trust_access_policy.%[1]s_p1.id
      decision = cloudflare_zero_trust_access_policy.%[1]s_p1.decision
      name = cloudflare_zero_trust_access_policy.%[1]s_p1.name
      include = [{
	email = { email = cloudflare_zero_trust_access_policy.%[1]s_p1.include.0.email.email }
      }]
    },
    {
      id = cloudflare_zero_trust_access_policy.%[1]s_p2.id
      decision = cloudflare_zero_trust_access_policy.%[1]s_p2.decision
      name = cloudflare_zero_trust_access_policy.%[1]s_p2.name
      include = [{
	ip = { ip = cloudflare_zero_trust_access_policy.%[1]s_p2.include.0.ip.ip }
      }]
    }
  ]
}
