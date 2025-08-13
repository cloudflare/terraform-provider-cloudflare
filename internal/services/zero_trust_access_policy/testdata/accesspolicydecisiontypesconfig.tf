resource "cloudflare_zero_trust_access_policy" "%[1]s_deny" {
  name           = "%[1]s-deny"
  account_id     = "%[3]s"
  decision       = "deny"
  include = [{
    email = {
      email = "blocked@example.com"
    }
  }]
}

resource "cloudflare_zero_trust_access_policy" "%[1]s_bypass" {
  name           = "%[1]s-bypass"
  account_id     = "%[3]s"
  decision       = "bypass"
  include = [{
    ip = {
      ip = "127.0.0.1/32"
    }
  }]
}