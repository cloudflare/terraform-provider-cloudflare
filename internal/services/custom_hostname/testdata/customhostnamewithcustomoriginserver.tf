
resource "cloudflare_custom_hostname" "%[2]s" {
  zone_id = "%[1]s"
  hostname = "%[2]s.%[3]s"
  custom_origin_server = "origin.%[2]s.terraform.cfapi.net"
	custom_origin_sni = "origin.%[2]s.terraform.cfapi.net"
  ssl = {
  method = "txt"
}
}

resource "cloudflare_record" "%[2]s" {
  zone_id = "%[1]s"
  name    = "origin.%[2]s.terraform.cfapi.net"
  value   = "example.com"
  type    = "CNAME"
  ttl     = 3600
}