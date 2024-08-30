
resource "cloudflare_zone" "%[1]s" {
	account = {
    id = "%[6]s"
  }
	name = "%[2]s"
	type = "%[7]s"
	vanity_name_servers = ["ns1.%[2]s", "ns2.%[2]s"]
}
