data "cloudflare_account_api_token_permission_groups" "example_account_api_token_permission_groups" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name = "Account%20Settings%20Write"
  scope = "com.cloudflare.api.account.zone"
}
