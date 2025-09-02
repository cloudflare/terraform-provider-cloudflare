resource "cloudflare_zone_subscription" "%[1]s" {
  zone_id = "%[2]s"
  
  # Specify the current rate plan to avoid changing it
  rate_plan = {
    id = "enterprise"
  }
}