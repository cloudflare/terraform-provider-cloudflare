resource "cloudflare_zero_trust_device_ip_profile" "example_zero_trust_device_ip_profile" {
  account_id = "account_id"
  match = "identity.email == \"test@cloudflare.com\""
  name = "IPv4 Cloudflare Source IPs"
  precedence = 100
  subnet_id = "b70ff985-a4ef-4643-bbbc-4a0ed4fc8415"
  description = "example comment"
  enabled = true
}
