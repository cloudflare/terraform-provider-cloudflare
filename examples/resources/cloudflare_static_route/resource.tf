resource "cloudflare_static_route" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  description = "New route for new prefix 192.0.2.0/24"
  prefix      = "192.0.2.0/24"
  nexthop     = "10.0.0.0"
  priority    = 100
  weight      = 10
  colo_names = [
    "den01"
  ]
  colo_regions = [
    "APAC"
  ]
}
