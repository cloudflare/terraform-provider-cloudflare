resource "cloudflare_zero_trust_gateway_pacfile" "example_zero_trust_gateway_pacfile" {
  account_id = "699d98642c564d2e855e9661899b7252"
  contents = "function FindProxyForURL(url, host) { return \"DIRECT\"; }"
  name = "Devops team"
  description = "PAC file for Devops team"
  slug = "pac_devops"
}
