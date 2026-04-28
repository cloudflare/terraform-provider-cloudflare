resource "cloudflare_user_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s-updated"
}
