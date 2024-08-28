
resource "cloudflare_zone" "%[1]s" {
	account = {
    id = "%[6]s"
  }
	name = "%[2]s"
	plan = {
	  id = "%[5]s"
	}
	type = "%[7]s"
}
