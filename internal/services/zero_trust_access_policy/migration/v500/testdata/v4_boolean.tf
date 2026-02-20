resource "cloudflare_access_policy" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  decision   = "allow"
  include {
    %[3]s = true
  }
}
