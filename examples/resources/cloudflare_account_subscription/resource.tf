resource "cloudflare_account_subscription" "example_account_subscription" {
  zone_id = "zone_id"
  frequency = "monthly"
  rate_plan = {
    id = "free"
    currency = "USD"
    externally_managed = false
    is_contract = false
    public_name = "Business Plan"
    scope = "zone"
    sets = ["string"]
  }
}
