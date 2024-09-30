
resource "cloudflare_zero_trust_access_identity_provider" "%[2]s" {
  account_id = "%[1]s"
  name = "%[2]s-updated"
  type = "github"
  config = {
  client_id = "test"
    client_secret = "secret"
}
}