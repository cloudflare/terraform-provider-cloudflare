variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_request_firewall_managed"
  kind    = "zone"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "execute"
      action_parameters = {
        id = "00000000000000000000000000000000"
      }
    }
  ]
}
