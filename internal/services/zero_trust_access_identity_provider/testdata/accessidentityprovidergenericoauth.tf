resource "cloudflare_zero_trust_access_identity_provider" "%[2]s" {
  account_id = "%[1]s"
  name = "%[2]s"
  type = "oauth2"
  config = {
    client_id = "test"
    client_secret = "secret"
    auth_url = "https://example.com/auth"
    token_url = "https://example.com/token"
    certs_url = "https://example.com/certs"
    scopes = ["openid", "profile", "email"]
    pkce_enabled = true
  }
}