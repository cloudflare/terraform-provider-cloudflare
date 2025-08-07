resource "cloudflare_zero_trust_access_identity_provider" "%[2]s" {
  account_id = "%[1]s"
  name = "%[2]s"
  type = "okta"
  config = {
    client_id = "test"
    client_secret = "secret"
    okta_account = "example.okta.com"
    authorization_server_id = "default"
  }
}