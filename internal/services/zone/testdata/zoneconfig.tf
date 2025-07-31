
resource "cloudflare_zone" "%[1]s" {
	account = {
    id = "%[3]s"
  }
	name = "%[2]s"
}
