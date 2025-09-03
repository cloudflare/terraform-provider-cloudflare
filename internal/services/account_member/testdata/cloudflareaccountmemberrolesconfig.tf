resource "cloudflare_account_member" "%[1]s" {
  account_id = "%[3]s"
  email      = "%[2]s"
  roles      = ["%[4]s"]
  status     = "pending"
}
