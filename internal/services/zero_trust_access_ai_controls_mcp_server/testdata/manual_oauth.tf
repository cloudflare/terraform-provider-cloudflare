resource "cloudflare_zero_trust_access_ai_controls_mcp_server" "tf-test-manual" {
  account_id = %[1]q
  id         = "tf-test-manual"
  auth_type  = "oauth"
  hostname   = "https://docs.mcp.cloudflare.com/mcp"
  name       = %[2]q

  auth_credentials = jsonencode({
    auth_mode = "manual"
    config = {
      authorization_endpoint = "https://auth.example.com/oauth/authorize"
      token_endpoint         = "https://auth.example.com/oauth/token"
    }
    registration_info = {
      client_id                  = "terraform-acceptance-test"
      token_endpoint_auth_method = "client_secret_basic"
      scope                      = "read"
    }
  })
  client_secret = %[3]q

  is_shared_oauth_callback_enabled = true
}
