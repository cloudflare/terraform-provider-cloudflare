resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  name = "%[1]s"
  account_id = "%[2]s"
  description = "desc"
  precedence = 12304
  action = "resolve"
  filters = ["dns_resolver"]
  traffic = "any(dns.domains[*] == \"example.com\")"
  rule_settings = {
    dns_resolvers = {
      ipv4 = [{
        ip = "2.2.2.2"
        port = 5053
      }]
      ipv6 = [{
        ip = "2001:DB8::"
        port = 5053
      }]
    }
  }
}
