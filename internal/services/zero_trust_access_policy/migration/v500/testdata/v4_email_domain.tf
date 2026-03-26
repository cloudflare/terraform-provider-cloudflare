resource "cloudflare_access_policy" {
  account_id = "%[2]s"
  name       = "%[1]s"
  decision   = "allow"

  include {
    email_domain = ["%[3]s"]
  }
}
