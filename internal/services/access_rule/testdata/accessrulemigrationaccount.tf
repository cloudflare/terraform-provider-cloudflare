resource "cloudflare_access_rule" "%[2]s" {
  account_id = "%[1]s"
  mode       = "block"
  notes      = "Block malicious IP"

  configuration {
    target = "ip"
    value  = "198.51.100.50"
  }
}
