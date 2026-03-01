resource "cloudflare_spectrum_application" "%s" {
  zone_id   = "%s"
  protocol  = "tcp/3306"
  dns {
    type = "CNAME"
    name = "%s.%s"
  }
  origin_port_range {
    start = 3306
    end   = 3310
  }
  origin_direct = ["tcp://128.66.0.1:23"]
}
