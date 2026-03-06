resource "cloudflare_account_member" "%s" {
  account_id    = "%s"
  email_address = "%s"
  status        = "accepted"
  role_ids      = ["%s"]
}
