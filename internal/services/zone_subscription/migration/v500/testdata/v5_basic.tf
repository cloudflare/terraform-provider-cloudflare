resource "cloudflare_zone_subscription" "%s" {
  zone_id = "%s"
  # frequency = "monthly"
  rate_plan = {
    id = "free"
  }
}
