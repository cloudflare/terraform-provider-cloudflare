
resource "cloudflare_zone" "%[1]s" {
	account = {
    id = "%[6]s"
  }
	name = "%[2]s"
}


resource "cloudflare_zone_subscription" "%[1]s" {
  zone_id = cloudflare_zone.%[1]s.id
  rate_plan = {
    id = "%[5]s"
  }
}
