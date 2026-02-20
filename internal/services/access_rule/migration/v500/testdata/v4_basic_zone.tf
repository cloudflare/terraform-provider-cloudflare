resource "cloudflare_access_rule" "%s" {
  zone_id = "%s"
  mode    = "block"
  notes   = "Test access rule"

  configuration {
    target = "ip"
    value  = "%s"
  }
}
