resource "cloudflare_zero_trust_access_identity_provider" "%[1]s" {
  %[2]s_id = "%[3]s"
  name     = "%[1]s"
  type     = "onetimepin"
  config = {
    client_name = "test_name"
  }
  scim_config = {
    enabled                  = false
    group_member_deprovision = false
    identity_update_behavior = "no_action"
    seat_deprovision         = false
    user_deprovision         = false
  }
}
