resource "cloudflare_list" "%[1]s" {
  account_id = "%[4]s"
  name = "%[2]s"
  description = "%[3]s"
  kind = "ip"
  items = [
    {
      ip = "1.1.1.1"
    },
    {
      ip = "1.1.1.2"
    },
  ]
}

data "cloudflare_list" "%[1]s" {
  account_id = "%[4]s"
  list_id = cloudflare_list.%[1]s.id
  search = "%[6]s"
}