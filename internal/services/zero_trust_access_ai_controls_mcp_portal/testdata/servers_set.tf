resource "cloudflare_zero_trust_access_ai_controls_mcp_server" "a" {
  account_id = %[1]q
  id         = "tf-test-srv-a"
  name       = "tf-test-srv-a"
  hostname   = "https://tf-test-srv-a.%[2]s/mcp"
  auth_type  = "bearer"
}

resource "cloudflare_zero_trust_access_ai_controls_mcp_server" "b" {
  account_id = %[1]q
  id         = "tf-test-srv-b"
  name       = "tf-test-srv-b"
  hostname   = "https://tf-test-srv-b.%[2]s/mcp"
  auth_type  = "bearer"
}

resource "cloudflare_zero_trust_access_ai_controls_mcp_portal" "tf-test" {
  account_id = %[1]q
  id         = "tf-test"
  hostname   = %[2]q
  name       = %[3]q

  servers = [
    { server_id = cloudflare_zero_trust_access_ai_controls_mcp_server.a.id },
    { server_id = cloudflare_zero_trust_access_ai_controls_mcp_server.b.id },
  ]
}
