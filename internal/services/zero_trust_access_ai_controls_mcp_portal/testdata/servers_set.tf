resource "cloudflare_zero_trust_access_ai_controls_mcp_server" "a" {
  account_id = %[1]q
  id         = "%[3]s-srv-a"
  name       = "%[3]s-srv-a"
  hostname   = "https://%[3]s-srv-a.%[2]s/mcp"
  auth_type  = "bearer"
}

resource "cloudflare_zero_trust_access_ai_controls_mcp_server" "b" {
  account_id = %[1]q
  id         = "%[3]s-srv-b"
  name       = "%[3]s-srv-b"
  hostname   = "https://%[3]s-srv-b.%[2]s/mcp"
  auth_type  = "bearer"
}

resource "cloudflare_zero_trust_access_ai_controls_mcp_portal" "tf-test" {
  account_id = %[1]q
  id         = %[3]q
  hostname   = %[2]q
  name       = %[3]q

  servers = [
    { server_id = cloudflare_zero_trust_access_ai_controls_mcp_server.a.id },
    { server_id = cloudflare_zero_trust_access_ai_controls_mcp_server.b.id },
  ]
}
