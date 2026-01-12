resource "cloudflare_access_rule" "%[2]s" {
  account_id = "%[1]s"
  mode       = "whitelist"

  configuration {
    target = "ip_range"
    value  = "192.0.2.0/24"
  }
}
