resource "cloudflare_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[3]s"
  kind        = "ip"
  description = "%[4]s"
}

resource "cloudflare_list_item" "%[1]s_item" {
  account_id = "%[2]s"
  list_id    = cloudflare_list.%[1]s.id
  ip         = "192.0.2.1"
  comment    = "Test IP item"
}
