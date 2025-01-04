
  resource "cloudflare_account_member" "%[1]s" {
	account_id = "%[3]s"
    email = "%[2]s"
    roles = [ "05784afa30c1afe1440e79d9351c7430" ]
  }
