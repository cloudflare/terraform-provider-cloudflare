resource "cloudflare_r2_bucket" "%s" {
  account_id    = "%s"
  name          = "%s"
  location      = "WNAM"
  jurisdiction  = "default"
  storage_class = "Standard"
}
