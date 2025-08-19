variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id     = var.zone_id
  name        = "My ruleset"
  description = ""
  phase       = "http_request_firewall_custom"
  kind        = "zone"
}

data "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  id      = cloudflare_ruleset.my_ruleset.id
}
