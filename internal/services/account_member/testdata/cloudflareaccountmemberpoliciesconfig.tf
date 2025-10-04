data "cloudflare_resource_groups" "example_resource_groups" {
  account_id = "%[1]s"
  name       = "com.cloudflare.api.account.%[1]s"
}

resource "cloudflare_account_member" "%[2]s" {
  account_id = "%[3]s"
  email      = "%[4]s"
  status     = "pending"
  policies = [{
    access = "allow"
    permission_groups = [{
      id = "%[5]s"
    }]
    resource_groups = [{
      id = data.cloudflare_resource_groups.example_resource_groups.result[0].id
    }]
  }]
}

