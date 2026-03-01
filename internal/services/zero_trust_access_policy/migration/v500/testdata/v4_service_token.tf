resource "cloudflare_access_policy" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  decision   = "allow"

  include {
    any_valid_service_token = true
  }
}
