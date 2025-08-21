
resource "cloudflare_zone" "%[1]s" {
	account = {
    id = "%[2]s"
  }
	name = "%[3]s"
	type = "%[4]s"
}
