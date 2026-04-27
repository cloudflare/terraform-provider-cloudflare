resource "cloudflare_user_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
}

data "cloudflare_user_groups" "%[1]s" {
  account_id = "%[2]s"
  depends_on = [cloudflare_user_group.%[1]s]
}

data "cloudflare_user_group" "%[1]s" {
  account_id    = "%[2]s"
  user_group_id = cloudflare_user_group.%[1]s.id
}
