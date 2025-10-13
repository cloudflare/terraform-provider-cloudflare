resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Egress policy"
  precedence  = 12403
  action      = "egress"
  filters     = ["egress"]
  traffic     = "net.dst.port in {443 80}"

  rule_settings = {
    egress = {
      ipv6 = "2001:db8::/64"
    }
  }
}