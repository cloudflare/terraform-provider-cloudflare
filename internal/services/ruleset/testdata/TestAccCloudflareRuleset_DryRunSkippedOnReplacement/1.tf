variable "account_id" {}
variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset_moving_to_the_zone" {
  account_id = var.account_id
  name       = "My ruleset moving to the zone"
  kind       = "custom"
  phase      = "http_request_firewall_custom"

  rules = [{
    action      = "block"
    expression  = "http.host eq \"example.com\""
    description = "Block example.com"
  }]
}

resource "cloudflare_ruleset" "my_ruleset_moving_to_the_account" {
  zone_id = var.zone_id
  name    = "My ruleset moving to the account"
  kind    = "custom"
  phase   = "http_request_firewall_custom"

  rules = [{
    action      = "block"
    expression  = "http.host eq \"example.com\""
    description = "Block example.com"
  }]
}
