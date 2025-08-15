
  resource "cloudflare_list" "%[2]s" {
    account_id          = "%[4]s"
    name                = "%[2]s"
    description         = "list named %[2]s"
    kind                = "hostname"
  }

  resource "cloudflare_list_item" "%[1]s_1" {
    account_id = "%[4]s"
  	list_id    = cloudflare_list.%[2]s.id
  	hostname         = {
      url_hostname = "a.example.com"
    }
  	comment    = "%[3]s"
  }

  resource "cloudflare_list_item" "%[1]s_2" {
    account_id = "%[4]s"
  	list_id    = cloudflare_list.%[2]s.id
  	hostname         = {
      url_hostname = "example.com"
    }
  	comment    = "%[3]s"
  }
