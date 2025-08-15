variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id     = var.zone_id
  name        = "My ruleset"
  description = "My updated ruleset description"
  phase       = "http_request_firewall_custom"
  kind        = "zone"
}
