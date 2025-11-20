
data "cloudflare_dns_record" "%[1]s" {
  zone_id = "%[2]s"
  type = "TXT"
  hostname = cloudflare_dns_record.%[1]s.hostname
}
resource "cloudflare_dns_record" "%[1]s" {
	zone_id = "%[2]s"
	type = "TXT"
	name = "%[1]s.%[3]s"
	content = "i am a text record"
}
