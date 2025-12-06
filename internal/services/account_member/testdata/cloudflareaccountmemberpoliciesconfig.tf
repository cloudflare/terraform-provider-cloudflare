data "cloudflare_resource_groups" "example_resource_groups" {
  account_id = "%[1]s"
  name       = "com.cloudflare.api.account.%[1]s"
}

resource "cloudflare_account_member" "test_member" {
  account_id = "%[1]s"
  email      = "%[2]s"
  status     = "pending"
  policies = [{
    access = "allow"
    permission_groups = [{
      id = "%[3]s"
    }]
    resource_groups = [{
      id = data.cloudflare_resource_groups.example_resource_groups.result[0].id
    }]
  }]
}

