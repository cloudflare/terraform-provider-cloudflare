resource "cloudflare_zone_subscription" "%[1]s" {
  zone_id = "%[2]s"

  rate_plan = {
    id = "free"
    currency = "USD"
    externally_managed = false
    is_contract = false
    public_name = "Cloudflare Free Plan"
    scope = "zone"
  }
}