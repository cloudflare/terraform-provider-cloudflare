variable "zone_id" {}

resource "terraform_data" "dependency" {
  input = "Managed by Terraform"
}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id     = var.zone_id
  name        = "My ruleset"
  description = terraform_data.dependency.output
  phase       = "http_request_firewall_managed"
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
