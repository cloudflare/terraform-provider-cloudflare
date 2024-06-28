data "cloudflare_account_roles" "account_roles" {
  account_id = "f037e56e89293a057740de681ac9abbe"
}

locals {
  roles_by_name = {
    for role in data.cloudflare_account_roles.account_roles.roles :
    role.name => role
  }
}

resource "cloudflare_account_member" "member" {
  account_id    = "f037e56e89293a057740de681ac9abbe"
  email_address = "user@example.com"
  role_ids = [
    local.roles_by_name["Administrator"].id
  ]
}
