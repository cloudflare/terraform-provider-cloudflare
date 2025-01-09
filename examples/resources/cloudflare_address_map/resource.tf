resource "cloudflare_address_map" "example_address_map" {
  account_id = "258def64c72dae45f3e4c8516e2111f2"
  description = "My Ecommerce zones"
  enabled = true
  ips = ["192.0.2.1"]
  memberships = [{
    can_delete = true
    created_at = "2014-01-01T05:20:00.12345Z"
    identifier = "023e105f4ecef8ad9ca31a8372d0c353"
    kind = "zone"
  }]
}
