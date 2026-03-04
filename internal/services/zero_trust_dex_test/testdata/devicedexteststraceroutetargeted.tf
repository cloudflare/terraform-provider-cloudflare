resource "cloudflare_zero_trust_dex_rule" "%[5]s" {
  account_id  = "%[2]s"
  name        = "%[5]s"
  match       = "%[6]s"
  description = "%[7]s"
}

resource "cloudflare_zero_trust_dex_test" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "%[4]s"
  interval    = "0h30m0s"
  enabled     = true
  data = {
    host   = "%[3]s"
    kind   = "traceroute"
  }
  target_policies = [
    {
      id = cloudflare_zero_trust_dex_rule.%[5]s.id
    }
  ]
}
