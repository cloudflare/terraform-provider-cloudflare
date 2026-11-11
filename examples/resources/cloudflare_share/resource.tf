resource "cloudflare_share" "example_share" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name = "My Shared WAF Managed Rule"
  excluded_recipients = [{
    account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  }]
  recipients = [{
    organization_id = "023e105f4ecef8ad9ca31a8372d0c353"
    recipient_account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  }]
  resources = [{
    meta = jsonencode({})
    resource_account_id = "023e105f4ecef8ad9ca31a8372d0c353"
    resource_id = "0a679032-bbb2-4b7e-9acc-dd5c5cb27a17"
    resource_type = "custom-ruleset"
  }]
  tags = {
    foo = "string"
  }
}
