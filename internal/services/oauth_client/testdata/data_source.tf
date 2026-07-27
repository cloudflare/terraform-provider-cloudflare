resource "cloudflare_oauth_client" "%[1]s" {
  account_id = "%[2]s"
  client_name = "%[1]s"
  grant_types = ["authorization_code"]
  redirect_uris = ["https://example.com/callback"]
  response_types = ["code"]
  scopes = ["user-details.read", "teams.read"]
  token_endpoint_auth_method = "none"
}

data "cloudflare_oauth_client" "%[1]s" {
  account_id = "%[2]s"
  oauth_client_id = cloudflare_oauth_client.%[1]s.client_id
}

data "cloudflare_oauth_clients" "%[1]s" {
  account_id = "%[2]s"
  depends_on = [cloudflare_oauth_client.%[1]s]
}
