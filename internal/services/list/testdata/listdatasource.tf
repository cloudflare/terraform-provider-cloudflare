resource "cloudflare_list" "%[1]s" {
  account_id = "%[2]s"
  name = "%[3]s"
  description = "%[4]s"
  kind = "%[5]s"
}

data "cloudflare_list" "%[1]s" {
  account_id = "%[2]s"
  list_id    = cloudflare_list.%[1]s.id
}
