resource "cloudflare_account_subscription" "example_account_subscription" {
  account_id = "account_id"
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
