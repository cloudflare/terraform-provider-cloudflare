
resource "cloudflare_zone" "%[1]s" {
	account = {
    id = "%[2]s"
  }
	name = "%[3]s"
	type = "%[4]s"
	vanity_name_servers = ["ns1.%[3]s", "ns2.%[3]s"]
}
