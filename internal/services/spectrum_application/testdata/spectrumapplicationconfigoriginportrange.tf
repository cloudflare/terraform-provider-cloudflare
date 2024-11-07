
resource "cloudflare_record" "%[3]s" {
	zone_id = "%[1]s"
	name    = "%[3]s.origin"
	value   = "example.com"
	type    = "CNAME"
	ttl     = 3600
}

resource "cloudflare_spectrum_application" "%[3]s" {
  depends_on = ["cloudflare_record.%[3]s"]

  zone_id  = "%[1]s"
  protocol = "tcp/22-23"

  dns = {
  type = "CNAME"
    name = "%[3]s.%[2]s"
}

  origin_dns = {
  name = "%[3]s.origin.%[2]s"
}
  origin_port_range = [{
    start = 2022
    end   = 2023
  }]

  edge_ips = {
  type = "dynamic"
	connectivity = "all"
}
}