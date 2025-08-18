variable "account_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  account_id = var.account_id
  name       = "My ruleset"
  phase      = "http_request_firewall_custom"
  kind       = "custom"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "block"
      exposed_credential_check = {
        username_expression = "lookup_json_string(http.request.body.raw, \"username\")"
        password_expression = "lookup_json_string(http.request.body.raw, \"password\")"
      }
    }
  ]
}

data "cloudflare_ruleset" "my_ruleset" {
  account_id = var.account_id
  id         = cloudflare_ruleset.my_ruleset.id
}
