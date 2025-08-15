resource "cloudflare_zero_trust_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"

  include = [
    {
      gsuite = {
        email                = "admin@example.com"
        identity_provider_id = "gsuite-idp-id"
      }
    }
  ]
}