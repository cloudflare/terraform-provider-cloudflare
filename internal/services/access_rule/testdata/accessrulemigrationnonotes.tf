resource "cloudflare_access_rule" "%[2]s" {
  account_id = "%[1]s"
  mode       = "whitelist"

  configuration {
    target = "ip_range"
    value  = "198.51.100.0/24"
  }
}
