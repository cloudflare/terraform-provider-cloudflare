resource "cloudflare_zero_trust_device_custom_profile" "%[1]s" {
  account_id               = "%[2]s"
  name                     = "%[1]s"
  match                    = "os.version == \"10.15\""
  precedence               = %[3]d
  enabled                  = true
  description              = "Complete custom device profile with all settings"
  allow_mode_switch        = true
  allow_updates            = true
  allowed_to_leave         = false
  auto_connect             = 60
  captive_portal           = 300
  disable_auto_fallback    = true
  exclude_office_ips       = true
  switch_locked            = true
  tunnel_protocol          = "wireguard"
  support_url              = "https://support.example.com"
  lan_allow_minutes        = 30
  lan_allow_subnet_size    = 24
  
  service_mode_v2 = {
    mode = "warp"
  }
}
