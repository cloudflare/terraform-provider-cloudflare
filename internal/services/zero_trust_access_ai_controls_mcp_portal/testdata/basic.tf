resource "cloudflare_zero_trust_access_ai_controls_mcp_portal" "tf-test" {
  account_id = %[1]q
  hostname   = %[2]q
  name       = %[3]q
  id         = "tf-test"
}
