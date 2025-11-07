resource "cloudflare_zero_trust_access_ai_controls_mcp_server" "example_zero_trust_access_ai_controls_mcp_server" {
  account_id = "a86a8f5c339544d7bdc89926de14fb8c"
  id = "my-mcp-server"
  auth_type = "unauthenticated"
  hostname = "https://exmaple.com/mcp"
  name = "My MCP Server"
  auth_credentials = "auth_credentials"
  description = "This is one remote mcp server"
}
