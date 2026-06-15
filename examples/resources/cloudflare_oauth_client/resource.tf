resource "cloudflare_oauth_client" "example_oauth_client" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  client_name = "My OAuth App"
  grant_types = ["authorization_code", "refresh_token"]
  redirect_uris = ["https://example.com/callback"]
  response_types = ["code"]
  scopes = ["account.read"]
  token_endpoint_auth_method = "client_secret_post"
  allowed_cors_origins = ["https://example.com"]
  client_uri = "https://example.com"
  logo_uri = "https://example.com/logo.png"
  policy_uri = "https://example.com/privacy"
  post_logout_redirect_uris = ["https://example.com/logout"]
  tos_uri = "https://example.com/tos"
}
