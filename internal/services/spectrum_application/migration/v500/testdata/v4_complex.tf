resource "cloudflare_spectrum_application" "%s" {
  zone_id   = "%s"
  protocol  = "tcp/443"
  dns {
    type = "CNAME"
    name = "%s.%s"
  }
  edge_ips {
    type         = "dynamic"
    connectivity = "all"
  }
  origin_direct = ["tcp://128.66.0.3:443"]
  tls              = "flexible"
  argo_smart_routing = true
  proxy_protocol   = "v1"
  ip_firewall      = true
  traffic_type     = "direct"
}
