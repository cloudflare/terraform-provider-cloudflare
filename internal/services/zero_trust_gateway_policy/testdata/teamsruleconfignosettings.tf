
resource "cloudflare_teams_rule" "%[1]s" {
  name = "%[1]s"
  account_id = "%[2]s"
  description = "desc"
  precedence = 12302
  action = "block"
  filters = ["dns"]
  traffic = "any(dns.domains[*] == \"example.com\")"
}
