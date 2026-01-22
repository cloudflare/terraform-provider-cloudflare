resource "cloudflare_account_member" "test_member" {
  account_id = "%[1]s"
  email      = "%[2]s"
  status     = "pending"
  roles = [
    "doesn't matter"
  ]
  policies = [{
    access = "allow"
    resource_groups = [{
      id : "doesn't matter"
    }]
    permission_groups = [{
      id : "doesn't matter"
    }]
  }]
}

