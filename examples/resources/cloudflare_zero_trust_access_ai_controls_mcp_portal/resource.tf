resource "cloudflare_zero_trust_access_ai_controls_mcp_portal" "example_zero_trust_access_ai_controls_mcp_portal" {
  account_id = "a86a8f5c339544d7bdc89926de14fb8c"
  id = "my-mcp-portal"
  hostname = "exmaple.com"
  name = "My MCP Portal"
  description = "This is my custom MCP Portal"
  secure_web_gateway = false
  servers = [{
    server_id = "my-mcp-server"
    default_disabled = true
    on_behalf = true
    updated_prompts = [{
      name = "name"
      description = "description"
      enabled = true
    }]
    updated_tools = [{
      name = "name"
      description = "description"
      enabled = true
    }]
  }]
}
