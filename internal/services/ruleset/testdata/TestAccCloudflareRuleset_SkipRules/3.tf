variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_request_firewall_managed"
  kind    = "zone"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "skip"
      action_parameters = {
        rules = {
          "4814384a9e5d4991b9815dcfc25d2f1f" = ["04116d14d7524986ba314d11c8a41e11"]
        }
        rulesets = ["4814384a9e5d4991b9815dcfc25d2f1f"]
      }
    }
  ]
}
