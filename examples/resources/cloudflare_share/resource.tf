resource "cloudflare_share" "example_share" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name = "My Shared WAF Managed Rule"
  recipients = [{
    organization_id = "023e105f4ecef8ad9ca31a8372d0c353"
    recipient_account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  }]
  resources = [{
    meta = {

    }
    resource_account_id = "023e105f4ecef8ad9ca31a8372d0c353"
    resource_id = "023e105f4ecef8ad9ca31a8372d0c353"
    resource_type = "custom-ruleset"
  }]
}
