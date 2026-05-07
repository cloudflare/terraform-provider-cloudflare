resource "cloudflare_user_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  policies = [
    {
      access            = "allow"
      permission_groups = [
        { id = "%[3]s" }
      ]
      resource_groups = [
        { id = "%[4]s" }
      ]
    }
  ]
}
