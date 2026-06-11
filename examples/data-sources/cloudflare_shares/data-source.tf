data "cloudflare_shares" "example_shares" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  include_recipient_counts = true
  include_resources = true
  kind = "sent"
  resource_types = ["custom-ruleset"]
  status = "active"
  tag = ["env=production"]
  target_type = "account"
}
