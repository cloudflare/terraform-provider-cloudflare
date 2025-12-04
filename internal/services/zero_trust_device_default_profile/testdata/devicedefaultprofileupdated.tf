resource "cloudflare_zero_trust_device_default_profile" "%[1]s" {
  account_id         = "%[2]s"
  allow_mode_switch  = false
  allow_updates      = false
  allowed_to_leave   = true
  auto_connect       = 120
  captive_portal     = 600
  switch_locked      = true
  support_url        = "https://updated-support.example.com"

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
