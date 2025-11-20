resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "My ruleset %[2]s"
  phase   = "http_request_firewall_custom"
  kind    = "zone"
}