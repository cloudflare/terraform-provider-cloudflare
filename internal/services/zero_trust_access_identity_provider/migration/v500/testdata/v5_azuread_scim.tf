resource "cloudflare_zero_trust_access_identity_provider" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "azureAD"
  config = {
    client_id     = "test"
    client_secret = "test"
    directory_id  = "directory"
  }
  scim_config = {
    enabled                   = true
    seat_deprovision          = true
    user_deprovision          = true
    identity_update_behavior  = "no_action"
  }
}
