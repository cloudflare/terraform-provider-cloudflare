resource "cloudflare_zero_trust_access_ai_controls_mcp_server" "example_zero_trust_access_ai_controls_mcp_server" {
  account_id = var.cloudflare_account_id
  id         = "github"
  auth_type  = "oauth"
  hostname   = "https://github-mcp.example.com/mcp"
  name       = "GitHub MCP Server"

  auth_credentials = jsonencode({
    auth_mode = "manual"
    config = {
      authorization_endpoint = "https://github.com/login/oauth/authorize"
      token_endpoint         = "https://github.com/login/oauth/access_token"
    }
    registration_info = {
      client_id                  = var.mcp_oauth_client_id
      token_endpoint_auth_method = "client_secret_basic"
      scope                      = "repo read:user"
    }
  })
  client_secret = var.mcp_oauth_client_secret

  # This lets the API fill in Cloudflare's shared callback URL.
  is_shared_oauth_callback_enabled = true
}

variable "cloudflare_account_id" {
  type = string
}

variable "mcp_oauth_client_id" {
  type = string
}

variable "mcp_oauth_client_secret" {
  type      = string
  sensitive = true
}
