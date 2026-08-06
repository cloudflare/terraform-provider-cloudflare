variable "account_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  account_id = var.account_id
  name       = "My ruleset"
  phase      = "http_request_firewall_managed"
  kind       = "root"
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
