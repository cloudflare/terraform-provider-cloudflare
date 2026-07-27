resource "cloudflare_account_subscription" "example_account_subscription" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  subscription_identifier = "506e3185e9c882d175a2d0cb0093d9f2"
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
