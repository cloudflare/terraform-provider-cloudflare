
resource "cloudflare_record" "%[2]s" {
	zone_id = "%[1]s"
	name = "%[2]s"
	value = "mail.terraform.cfapi.net"
	type = "MX"
	priority = 0
	proxied = false
	ttl = 300
}