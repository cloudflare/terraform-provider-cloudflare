resource "cloudflare_zero_trust_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"

  include = [
    {
      okta = {
        name                 = "test-okta-group"
        identity_provider_id = "okta-idp-id"
      }
    }
  ]
}