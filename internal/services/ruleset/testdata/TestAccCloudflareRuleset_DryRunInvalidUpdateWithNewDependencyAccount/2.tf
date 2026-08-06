variable "account_id" {}

resource "terraform_data" "dependency" {
  input = "Managed by Terraform"
}

resource "cloudflare_ruleset" "my_ruleset" {
  account_id  = var.account_id
  name        = "My ruleset"
  description = terraform_data.dependency.output
  phase       = "http_request_firewall_custom"
  kind        = "custom"
  rules = [
    {
      expression = "http.invalid_field eq \"x\""
      action     = "block"
    }
  ]
}
