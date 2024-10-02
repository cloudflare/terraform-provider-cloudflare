
  resource "cloudflare_account_member" "%[1]s" {
	account_id = "%[3]s"
    email_address = "%[2]s"
    role_ids = [ "05784afa30c1afe1440e79d9351c7430" ]
	status = "accepted"
  }