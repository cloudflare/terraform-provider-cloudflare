data "cloudflare_account_roles" "account_roles" {
    account_id = "goo"
}

locals {
  roles_by_name = {
    for role in data.cloudflare_account_roles.account_roles.roles :
      role.name => role
  }
}

resource "cloudflare_account_member" "member" {
  ...
  role_ids = [
    local.roles_by_name["Administrator"].id
  ]
}
