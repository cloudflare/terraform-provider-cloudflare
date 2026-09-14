resource "cloudflare_zero_trust_connectivity_settings" "example_zero_trust_connectivity_settings" {
  account_id = "699d98642c564d2e855e9661899b7252"
  icmp_proxy_enabled = true
  offramp_warp_enabled = true
}
