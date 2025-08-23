
  resource "cloudflare_list" "%[2]s" {
    account_id          = "%[4]s"
    name                = "%[2]s"
    description         = "list named %[2]s"
    kind                = "hostname"
  }

  resource "cloudflare_list_item" "%[1]s_1" {
    account_id = "%[4]s"
  	list_id    = cloudflare_list.%[2]s.id
  	comment    = "%[3]s"
  	hostname = {
  		url_hostname = "example.com"
  	}
  }

  resource "cloudflare_list_item" "%[1]s_2" {
    account_id = "%[4]s"
    list_id    = cloudflare_list.%[2]s.id
    comment    = "%[3]s"
    hostname = {
      url_hostname = "*.example.com"
      exclude_exact_hostname = true
    }
  }
