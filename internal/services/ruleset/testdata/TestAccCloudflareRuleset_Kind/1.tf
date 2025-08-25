variable "account_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  account_id = var.account_id
  name       = "My ruleset"
  phase      = "http_request_firewall_custom"
  kind       = "root"
}

data "cloudflare_ruleset" "my_ruleset" {
  account_id = var.account_id
  id         = cloudflare_ruleset.my_ruleset.id
}
