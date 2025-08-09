resource "cloudflare_spectrum_application" "%[3]s" {
  zone_id  = "%[1]s"
  protocol = "tcp/443"

  dns = {
    type = "CNAME"
    name = "%[3]s.%[2]s"
  }

  origin_direct = ["tcp://128.66.0.5:443"]
  origin_port   = 443
  
  # Test TLS configuration
  tls = "%[4]s"  # flexible, full, strict, or off
  
  edge_ips = {
    type = "dynamic"
    connectivity = "all"
  }
}
