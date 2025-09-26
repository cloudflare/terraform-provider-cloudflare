resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Safe search policy"
  precedence  = 12404
  action      = "safesearch"
  filters     = ["dns"]
  traffic     = "any(dns.domains[*] in {\"google.com\" \"bing.com\" \"duckduckgo.com\"})"
  enabled     = true
}