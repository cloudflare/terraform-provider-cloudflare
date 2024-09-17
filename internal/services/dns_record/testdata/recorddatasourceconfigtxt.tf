
data "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  type = "TXT"
  hostname = cloudflare_record.%[1]s.hostname
}
resource "cloudflare_record" "%[1]s" {
	zone_id = "%[2]s"
	type = "TXT"
	name = "%[1]s.%[3]s"
	value = "i am a text record"
}