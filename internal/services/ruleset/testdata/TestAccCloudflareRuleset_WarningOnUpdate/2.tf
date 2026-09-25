variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_request_firewall_custom"
  kind    = "custom"

  description = join("", [for i in range(410) : "0123456789"])

  rules = [{
    action     = "block"
    expression = "http.host eq \"never-matches.example\""
  }]
}
