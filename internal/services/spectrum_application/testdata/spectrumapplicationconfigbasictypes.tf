
resource "cloudflare_spectrum_application" "%[3]s" {
  zone_id  = "%[1]s"
  protocol = "%[4]s"

  dns = {
  type = "CNAME"
    name = "%[3]s.%[2]s"
}

  origin_direct = ["tcp://128.66.0.4:%[5]d"]
}