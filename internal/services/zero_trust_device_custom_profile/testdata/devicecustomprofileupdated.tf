resource "cloudflare_zero_trust_device_custom_profile" "%[1]s" {
  account_id               = "%[2]s"
  name                     = "%[1]s-updated"
  match                    = "os.version == \"10.15\""
  precedence               = %[3]d
  enabled                  = false
  description              = "Updated custom device profile"
  allow_mode_switch        = false
  allow_updates            = false
  allowed_to_leave         = true
  auto_connect             = 0
  captive_portal           = 180
  disable_auto_fallback    = false
  exclude_office_ips       = false
  switch_locked            = false
  tunnel_protocol          = "masque"
  support_url              = "https://help.example.com"
}
