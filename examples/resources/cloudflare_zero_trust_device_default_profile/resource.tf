resource "cloudflare_zero_trust_device_default_profile" "example_zero_trust_device_default_profile" {
  account_id = "699d98642c564d2e855e9661899b7252"
  allow_mode_switch = true
  allow_updates = true
  allowed_to_leave = true
  auto_connect = 0
  captive_portal = 180
  disable_auto_fallback = true
  exclude = [{
    address = "192.0.2.0/24"
    description = "Exclude testing domains from the tunnel"
  }]
  exclude_office_ips = true
  include = [{
    address = "192.0.2.0/24"
    description = "Include testing domains in the tunnel"
  }]
  lan_allow_minutes = 30
  lan_allow_subnet_size = 24
  register_interface_ip_with_dns = true
  sccm_vpn_boundary_support = false
  service_mode_v2 = {
    mode = "proxy"
    port = 3000
  }
  support_url = "https://1.1.1.1/help"
  switch_locked = true
  tunnel_protocol = "wireguard"
}
