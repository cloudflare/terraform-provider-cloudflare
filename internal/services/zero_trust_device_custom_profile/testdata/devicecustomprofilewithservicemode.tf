resource "cloudflare_zero_trust_device_custom_profile" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  match       = "os.version == \"10.15\""
  precedence  = %[3]d
  enabled     = true
  description = "Profile with service mode configuration"
  
  service_mode_v2 = {
    mode = "proxy"
    port = 3128
  }
}
