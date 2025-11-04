resource "cloudflare_zero_trust_access_ai_controls_mcp_server" "tf-test" {
  account_id = %[1]q
  hostname   = %[2]q
  name       = %[3]q
  auth_type  = "unauthenticated"
  id         = "tf-test"
}
