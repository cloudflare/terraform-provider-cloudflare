
data "cloudflare_dns_record" "%[1]s" {
  zone_id = "%[2]s"
  hostname = cloudflare_dns_record.%[1]s.hostname
}
resource "cloudflare_dns_record" "%[1]s" {
	zone_id = "%[2]s"
	type = "A"
	name = "%[1]s.%[3]s"
	content = "192.0.2.0"
}
