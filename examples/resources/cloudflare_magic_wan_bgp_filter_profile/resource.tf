resource "cloudflare_magic_wan_bgp_filter_profile" "example_magic_wan_bgp_filter_profile" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  match_action = "allow"
  name = "Allowed On-Prem Imports"
  targets = ["10.0.0.0/8{8,32}"]
  description = "Allowed corporate subnets from on-premises"
}
