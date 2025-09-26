resource "cloudflare_magic_wan_gre_tunnel" "example_magic_wan_gre_tunnel" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  cloudflare_gre_endpoint = "203.0.113.1"
  customer_gre_endpoint = "203.0.113.1"
  interface_address = "192.0.2.0/31"
  name = "GRE_1"
  automatic_return_routing = true
  bgp = {
    customer_asn = 0
    extra_prefixes = ["string"]
    md5_key = "md5_key"
  }
  description = "Tunnel for ISP X"
  health_check = {
    direction = "bidirectional"
    enabled = true
    rate = "low"
    target = {
      saved = "203.0.113.1"
    }
    type = "request"
  }
  interface_address6 = "2606:54c1:7:0:a9fe:12d2:1:200/127"
  mtu = 0
  ttl = 0
}
