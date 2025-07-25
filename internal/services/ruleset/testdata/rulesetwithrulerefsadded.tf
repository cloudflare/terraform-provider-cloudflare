
resource "cloudflare_ruleset" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[1]s"
  phase   = "http_request_firewall_custom"
  kind    = "zone"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1",
      action     = "block",
      ref        = "one",
    },
    {
      expression = "ip.src eq 3.3.3.3",
      action     = "block",
      ref        = "three",
    },
    {
      expression = "ip.src eq 2.2.2.2",
      action     = "block",
      ref        = "two",
    },
  ]
}
