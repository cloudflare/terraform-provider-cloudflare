
resource "cloudflare_zone" "%[1]s" {
	account = {
    id = "%[5]s"
  }
	name = "%[2]s"
}
