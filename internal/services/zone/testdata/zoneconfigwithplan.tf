
resource "cloudflare_zone" "%[1]s" {
	account = {
    id = "%[6]s"
  }
	name = "%[2]s"
	plan = "%[5]s"
}
