data "cloudflare_resource_groups" "example_resource_groups" {
  account_id = "%s"
}

resource "cloudflare_account_member" "%s" {
  account_id = "%s"
  email      = "%s"
  status     = "pending"
  policies = [{
    access = "allow"
    permission_groups = [{
      id = "%s"
    }]
    resource_groups = [{
      id = data.cloudflare_resource_groups.example_resource_groups.result[0].id
    }]
  }]
}
