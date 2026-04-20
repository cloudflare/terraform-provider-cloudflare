resource "cloudflare_account_member" "%s" {
  account_id = "%s"
  email      = "%s"
  status     = "pending"
  roles      = ["%s"]
}
