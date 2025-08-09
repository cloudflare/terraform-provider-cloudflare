resource "cloudflare_zero_trust_access_identity_provider" "%[2]s" {
  account_id = "%[1]s"
  name       = "%[2]s"
  type       = "azureAD"
  config = {
    client_id      = "test"
    client_secret  = "test"
    directory_id   = "directory"
    support_groups = true
    conditional_access_enabled = true
    prompt = "select_account"
  }
  scim_config = {
    enabled          = true
    seat_deprovision = true
    user_deprovision = true
    identity_update_behavior = "automatic"
  }
}