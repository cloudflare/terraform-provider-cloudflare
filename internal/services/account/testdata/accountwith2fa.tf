
  resource "cloudflare_account" "%[1]s" {
	  name = "%[2]s"
	  enforce_twofactor = true
  }