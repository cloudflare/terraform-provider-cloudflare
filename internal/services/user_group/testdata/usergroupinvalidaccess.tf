resource "cloudflare_user_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  policies = [
    {
      access = "invalid_action"
      permission_groups = [
        { id = "00000000000000000000000000000000" }
      ]
      resource_groups = [
        { id = "00000000000000000000000000000000" }
      ]
    }
  ]
}
