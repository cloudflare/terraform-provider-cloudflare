resource "cloudflare_zone" "%[3]s" {
  account = {
    id = "%[4]s"
  }
  name = "%[5]s"
  type = "full"
}

resource "cloudflare_zone_subscription" "%[1]s" {
  zone_id = cloudflare_zone.%[3]s.id

  rate_plan = {
    id = "%[2]s"
  }
}