resource "cloudflare_r2_bucket" "%s" {
  account_id    = "%s"
  name          = "%s"
  jurisdiction  = "default"
  storage_class = "Standard"
}
