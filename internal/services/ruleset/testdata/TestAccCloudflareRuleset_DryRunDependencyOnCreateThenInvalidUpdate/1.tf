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
}

resource "cloudflare_ruleset" "my_zone_ruleset" {
  zone_id     = var.zone_id
  name        = "My ruleset"
  description = terraform_data.dependency.output
  phase       = "ddos_l7"
  kind        = "zone"
}
