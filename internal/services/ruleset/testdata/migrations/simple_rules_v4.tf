resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "My ruleset %[2]s"
  phase   = "http_request_firewall_custom"
  kind    = "zone"
  rules {
    expression = "ip.src eq 1.1.1.1"
    action     = "block"
  }
  rules {
    expression = "ip.src eq 2.2.2.2"
    action     = "log"
  }
}