resource "cloudflare_access_group" "%s" {
  account_id = "%s"
  name       = "%s"

  include {
    everyone = true
  }
}
