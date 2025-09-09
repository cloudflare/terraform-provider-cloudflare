resource "cloudflare_dns_record" "%[3]s" {
	zone_id = "%[1]s"
	name    = "%[3]s.origin.%[2]s"
	content = "128.66.0.1"
	type    = "A"
	ttl     = 3600
}

resource "cloudflare_spectrum_application" "%[3]s" {
  depends_on = ["cloudflare_dns_record.%[3]s"]
  zone_id  = "%[1]s"
  protocol = "tcp/8080-8081"

  dns = {
    type = "CNAME"
    name = "%[3]s.%[2]s"
  }

  origin_dns = {
    name = "%[3]s.origin.%[2]s"
  }
  origin_port   = "8080-8081"

  edge_ips = {
    type = "dynamic"
  	connectivity = "all"
  }
}