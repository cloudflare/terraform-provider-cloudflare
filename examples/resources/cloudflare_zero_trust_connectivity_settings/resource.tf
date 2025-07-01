resource "cloudflare_zero_trust_connectivity_settings" "example" {
  account_id          = "f037e56e89293a057740de681ac9abbe"
  icmp_proxy_enabled  = true
  offram_warp_enabled = true
}
