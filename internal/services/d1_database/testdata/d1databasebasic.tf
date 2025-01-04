
  resource "cloudflare_d1_database" "%[1]s" {
    account_id = "%[2]s"
    name       = "%[1]s"
  }