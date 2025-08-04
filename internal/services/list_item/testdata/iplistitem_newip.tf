
  resource "cloudflare_list" "%[2]s" {
    account_id          = "%[4]s"
    name                = "%[2]s"
    description         = "list named %[2]s"
    kind                = "ip"
  }

  resource "cloudflare_list_item" "%[1]s" {
    account_id = "%[4]s"
  	list_id    = cloudflare_list.%[2]s.id
  	ip         = "%[5]s"
  	comment    = "%[3]s"
  }