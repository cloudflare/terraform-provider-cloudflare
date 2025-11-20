resource "cloudflare_zone_subscription" "%[1]s" {
  zone_id = "%[2]s"
  rate_plan = {
    id = "free"
  }
}
