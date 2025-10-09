resource "cloudflare_zero_trust_access_application" "%[1]s_mcp_server" {
  account_id       = "%[3]s"
  name             = "%[1]s_mcp_server"
  type             = "mcp"
  destinations = [
    {
      "type": "via_mcp_server_portal",
      "mcp_server_id": "%[1]s"
    }
  ]
}

resource "cloudflare_zero_trust_access_application" "%[1]s_mcp_portal" {
  account_id       = "%[3]s"
  name             = "%[1]s_mcp_portal"
  type             = "mcp_portal"
  session_duration = "24h"
  domain           = "%[1]s.%[2]s"
  self_hosted_domains = ["%[1]s.%[2]s"]
}