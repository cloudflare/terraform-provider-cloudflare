variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_request_firewall_custom"
  kind    = "zone"
  rules = [
    {
      description = "My rule description"
      expression  = "ip.src eq 1.1.1.1"
      action      = "block"
    }
  ]
}

data "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  id      = cloudflare_ruleset.my_ruleset.id
}
