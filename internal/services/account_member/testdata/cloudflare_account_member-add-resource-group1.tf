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
      id = "%[4]s"
    }]
  }]
}

