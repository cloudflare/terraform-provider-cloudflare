resource "cloudflare_spectrum_application" "%s" {
  zone_id  = "%s"
  protocol = "tcp/22"
  dns = {
    type = "CNAME"
    name = "%s.%s"
  }
  origin_direct = ["tcp://128.66.0.1:23"]
}
