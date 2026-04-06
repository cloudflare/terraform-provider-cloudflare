data "cloudflare_resource_groups" "all" {
  account_id = "%s"
  name       = "com.cloudflare.api.account.%s"
}

data "cloudflare_account_permission_groups" "all" {
  account_id = "%s"
}

locals {
  api_token_permissions_groups_map = {
    for perm in data.cloudflare_account_permission_groups.all.result :
    perm.name => perm.id
  }
}

resource "cloudflare_account_member" "%s" {
  account_id = "%s"
  email      = "%s"
  policies = [{
    access = "allow"
    resource_groups = [{
      id : data.cloudflare_resource_groups.all.result[0].id
    }]
    permission_groups = [{
      id : local.api_token_permissions_groups_map["Administrator"]
    }]
  }]
}
