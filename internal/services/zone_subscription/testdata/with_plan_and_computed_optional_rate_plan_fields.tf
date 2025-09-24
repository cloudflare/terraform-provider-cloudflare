resource "cloudflare_zone_subscription" "%[1]s" {
  zone_id = "%[2]s"

  rate_plan = {
    id = "enterprise"
    currency = "USD"
    externally_managed = false
    is_contract = true
    public_name = "Enterprise"
    scope = "zone"
  }
}