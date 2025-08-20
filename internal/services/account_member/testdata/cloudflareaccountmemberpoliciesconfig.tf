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
      // TODO: add resource groups id to make the test work, need to expose resource groups endpoint in terraform
      scope = {
        key = "com.cloudflare.api.account.%s"
        objects = [{
          key = "*"
        }]
      }
    }]
  }]
}
