resource "cloudflare_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[3]s"
  kind        = "ip"
  description = "%[4]s"
}

resource "cloudflare_list_item" "%[1]s_item1" {
  account_id = "%[2]s"
  list_id    = cloudflare_list.%[1]s.id
  ip         = "172.16.0.1"
  comment    = "Separate item 1"
}

resource "cloudflare_list_item" "%[1]s_item2" {
  account_id = "%[2]s"
  list_id    = cloudflare_list.%[1]s.id
  ip         = "172.16.0.2"
  comment    = "Separate item 2"
}