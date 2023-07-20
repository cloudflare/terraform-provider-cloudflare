resource "cloudflare_device_settings_policy" "developer_warp_policy" {
  account_id            = "f037e56e89293a057740de681ac9abbe"
  name                  = "Developers WARP settings policy"
  description           = "Developers WARP settings policy description"
  precedence            = 10
  match                 = "any(identity.groups.name[*] in {\"Developers\"})"
  default               = false
  enabled               = true
  allow_mode_switch     = true
  allow_updates         = true
  allowed_to_leave      = true
  auto_connect          = 0
  captive_portal        = 5
  disable_auto_fallback = true
  support_url           = "https://cloudflare.com"
  switch_locked         = true
  service_mode_v2_mode  = "warp"
  service_mode_v2_port  = 3000
  exclude_office_ips    = false
}
