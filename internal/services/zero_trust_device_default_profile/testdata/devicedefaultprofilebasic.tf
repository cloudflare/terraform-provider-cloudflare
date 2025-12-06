resource "cloudflare_zero_trust_device_default_profile" "%[1]s" {
  account_id         = "%[2]s"
  allow_mode_switch  = true
  allow_updates      = true
  allowed_to_leave   = false
  auto_connect       = 60
  captive_portal     = 300
  switch_locked      = false

  exclude = [
    {
      address     = "192.168.1.0/24"
      description = "Local network"
    }
  ]

  service_mode_v2 = {
    mode = "warp"
  }
}
