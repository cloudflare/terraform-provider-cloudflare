variable "account_id" {}
variable "zone_id" {}

resource "terraform_data" "dependency" {
  input = "my_rule_ref"
}

resource "cloudflare_ruleset" "my_account_ruleset" {
  account_id = var.account_id
  name       = "My ruleset"
  phase      = "http_request_firewall_managed"
  kind       = "root"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "execute"
      ref        = terraform_data.dependency.output
      action_parameters = {
        id = "00000000000000000000000000000000"
      }
    }
  ]
}

resource "cloudflare_ruleset" "my_zone_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_request_firewall_managed"
  kind    = "zone"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "execute"
      ref        = terraform_data.dependency.output
      action_parameters = {
        id = "00000000000000000000000000000000"
      }
    }
  ]
}
