variable "account_id" {}
variable "zone_id" {}

resource "terraform_data" "dependency" {
  input = "Managed by Terraform"
}

resource "cloudflare_ruleset" "my_account_ruleset" {
  account_id  = var.account_id
  name        = "My ruleset"
  description = terraform_data.dependency.output
  phase       = "ddos_l7"
  kind        = "root"
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

resource "cloudflare_ruleset" "my_zone_ruleset" {
  zone_id     = var.zone_id
  name        = "My ruleset"
  description = terraform_data.dependency.output
  phase       = "ddos_l7"
  kind        = "zone"
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
