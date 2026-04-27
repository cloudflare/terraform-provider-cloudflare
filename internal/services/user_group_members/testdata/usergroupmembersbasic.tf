resource "cloudflare_user_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
}

resource "cloudflare_user_group_members" "%[1]s" {
  account_id    = "%[2]s"
  user_group_id = cloudflare_user_group.%[1]s.id
  members = [
    { id = "%[3]s" }
  ]
}
