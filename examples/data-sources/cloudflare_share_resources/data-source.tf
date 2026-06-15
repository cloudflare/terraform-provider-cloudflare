data "cloudflare_share_resources" "example_share_resources" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  share_id = "3fd85f74b32742f1bff64a85009dda07"
  resource_type = "custom-ruleset"
  status = "active"
}
