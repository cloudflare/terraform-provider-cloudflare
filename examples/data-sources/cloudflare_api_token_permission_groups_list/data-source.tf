data "cloudflare_api_token_permission_groups_list" "example_api_token_permission_groups_list" {
  name = "Account%20Settings%20Write"
  scope = "com.cloudflare.api.account.zone"
}
