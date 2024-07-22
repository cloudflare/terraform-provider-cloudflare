
data "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  type = "MX"
  priority = 10
  hostname = cloudflare_record.%[1]s.hostname
}
resource "cloudflare_record" "%[1]s" {
	zone_id = "%[2]s"
	type = "MX"
	name = "%[1]s.%[3]s"
	value = "mx1.example.com"
	priority = 10
}
resource "cloudflare_record" "%[1]s_2" {
	zone_id = "%[2]s"
	type = "MX"
	name = "%[1]s.%[3]s"
	value = "mx1.example.com"
	priority = 20
}
