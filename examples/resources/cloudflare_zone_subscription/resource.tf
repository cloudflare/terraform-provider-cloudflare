resource "cloudflare_zone_subscription" "example_zone_subscription" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
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
