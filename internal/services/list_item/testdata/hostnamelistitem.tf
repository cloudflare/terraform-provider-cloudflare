
  # TODO: once cloudflare_list is working, uncomment the
  # resource
  #
  # resource "cloudflare_list" "%[2]s" {
  #   account_id          = "%[4]s"
  #   name                = "%[2]s"
  #   description         = "list named %[2]s"
  #   kind                = "hostname"
  # }

  resource "cloudflare_list_item" "%[1]s" {
    account_id = "%[4]s"
  	list_id    = "d6be21998bf74a1bb778b110a7b2864e" # cloudflare_list.%[2]s.id
  	comment    = "%[3]s"
  	hostname = {
  		url_hostname = "example.com"
  	}
  }
