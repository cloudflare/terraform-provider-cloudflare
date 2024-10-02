
  resource "cloudflare_r2_bucket" "%[1]s" {
    account_id = "%[2]s"
    name       = "%[1]s"
    location   = "ENAM"
  }
