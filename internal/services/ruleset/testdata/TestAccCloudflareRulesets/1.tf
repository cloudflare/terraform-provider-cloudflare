variable "account_id" {}

variable "zone_id" {}

resource "cloudflare_ruleset" "my_first_ruleset" {
  account_id = var.account_id
  name       = "My first ruleset"
  phase      = "http_request_firewall_custom"
  kind       = "root"
}

resource "cloudflare_ruleset" "my_second_ruleset" {
  account_id = var.account_id
  name       = "My second ruleset"
  phase      = "http_request_firewall_managed"
  kind       = "root"
}

data "cloudflare_rulesets" "my_zone_rulesets" {
  zone_id = var.zone_id
  depends_on = [
    cloudflare_ruleset.my_third_ruleset,
    cloudflare_ruleset.my_fourth_ruleset
  ]
}

resource "cloudflare_ruleset" "my_third_ruleset" {
  zone_id = var.zone_id
  name    = "My third ruleset"
  phase   = "http_request_firewall_custom"
  kind    = "zone"
}

resource "cloudflare_ruleset" "my_fourth_ruleset" {
  zone_id = var.zone_id
  name    = "My fourth ruleset"
  phase   = "http_request_firewall_managed"
  kind    = "zone"
}

data "cloudflare_rulesets" "my_account_rulesets" {
  account_id = var.account_id
  depends_on = [
    cloudflare_ruleset.my_first_ruleset,
    cloudflare_ruleset.my_second_ruleset
  ]
}
