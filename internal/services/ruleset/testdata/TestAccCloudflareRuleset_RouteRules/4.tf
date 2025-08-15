variable "zone_id" {}

variable "domain" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_request_origin"
  kind    = "zone"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "route"
      action_parameters = {
        sni = {
          value = var.domain
        }
      }
    }
  ]
}
