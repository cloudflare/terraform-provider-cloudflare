resource "cloudflare_zero_trust_device_subnet" "example_zero_trust_device_subnet" {
  account_id = "699d98642c564d2e855e9661899b7252"
  name = "IPv4 Cloudflare Source IPs"
  network = "100.64.0.0/12"
  comment = "example comment"
  is_default_network = true
}
