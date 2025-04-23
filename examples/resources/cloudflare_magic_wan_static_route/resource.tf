resource "cloudflare_magic_wan_static_route" "example_magic_wan_static_route" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  nexthop = "203.0.113.1"
  prefix = "192.0.2.0/24"
  priority = 0
  description = "New route for new prefix 203.0.113.1"
  scope = {
    colo_names = ["den01"]
    colo_regions = ["APAC"]
  }
  weight = 0
}
