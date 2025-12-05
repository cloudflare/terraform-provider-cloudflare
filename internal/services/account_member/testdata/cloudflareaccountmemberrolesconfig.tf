resource "cloudflare_account_member" "test_member" {
  account_id = "%[2]s"
  email      = "%[1]s"
  roles      = ["%[3]s"]
  status     = "pending"
}
