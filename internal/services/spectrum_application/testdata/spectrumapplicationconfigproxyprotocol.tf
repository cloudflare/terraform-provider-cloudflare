resource "cloudflare_spectrum_application" "%[3]s" {
  zone_id  = "%[1]s"
  protocol = "tcp/22"

  dns = {
    type = "CNAME"
    name = "%[3]s.%[2]s"
  }

  origin_direct = ["tcp://128.66.0.6:22"]
  origin_port   = 22
  
  # Test proxy protocol configuration
  proxy_protocol = "%[4]s"  # off, v1, v2, simple is not supported for TCP
  
  edge_ips = {
    type = "dynamic"
    connectivity = "all"
  }
}
