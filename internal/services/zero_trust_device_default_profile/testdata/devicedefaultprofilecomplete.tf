resource "cloudflare_zero_trust_device_default_profile" "%[1]s" {
  account_id                       = "%[2]s"
  allow_mode_switch                = true
  allow_updates                    = true
  allowed_to_leave                 = false
  auto_connect                     = 60
  captive_portal                   = 300
  disable_auto_fallback            = true
  exclude_office_ips               = true
  register_interface_ip_with_dns   = false
  sccm_vpn_boundary_support        = true
  switch_locked                    = true
  tunnel_protocol                  = "wireguard"
  support_url                      = "https://support.example.com"
  lan_allow_minutes                = 30
  lan_allow_subnet_size            = 24

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
