resource "cloudflare_oauth_client" "%[1]s" {
  account_id = "%[2]s"
  client_name = "%[3]s"
  grant_types = ["authorization_code"]
  redirect_uris = ["%[4]s"]
  response_types = ["code"]
  scopes = ["user-details.read", "teams.read"]
  token_endpoint_auth_method = "none"
}
