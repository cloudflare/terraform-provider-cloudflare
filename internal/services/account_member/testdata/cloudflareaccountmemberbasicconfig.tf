
data "cloudflare_resource_groups" "all" {
  account_id = "%[1]s"
  name       = "com.cloudflare.api.account.%[1]s"
}

data "cloudflare_account_permission_groups" "all" {
  account_id = "%[1]s"
}

locals {
  api_token_permissions_groups_map = {
    for perm in data.cloudflare_account_permission_groups.all.result :
    perm.name => perm.id
  }
}

resource "cloudflare_account_member" "test_member" {
  account_id = "%[1]s"
  email      = "%[2]s"
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
