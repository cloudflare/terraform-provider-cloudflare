resource "cloudflare_zero_trust_organization_trust_settings" "example_zero_trust_organization_trust_settings" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  grant_trust_to_all_accounts_in_org = false
  trusting_accounts = ["023e105f4ecef8ad9ca31a8372d0c353"]
}
