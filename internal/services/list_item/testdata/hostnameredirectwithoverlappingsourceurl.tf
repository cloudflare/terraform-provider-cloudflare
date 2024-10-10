
  # TODO: once cloudflare_list is working, uncomment the
  # resource
  #
  # resource "cloudflare_list" "%[2]s" {
  #   account_id          = "%[4]s"
  #   name                = "%[2]s"
  #   description         = "list named %[2]s"
  #   kind                = "redirect"
  # }

  resource "cloudflare_list_item" "%[1]s_1" {
    account_id = "%[4]s"
  	list_id    = "a2d365108e1840e99afc6cf4cc800869" # cloudflare_list.%[2]s.id
  	comment    = "%[3]s"
  	redirect = {
  		source_url = "www.site1.com/"
  		target_url = "https://example.com"
  		status_code = 301
  	}
  }

  resource "cloudflare_list_item" "%[1]s_2" {
    account_id = "%[4]s"
  	list_id    = "a2d365108e1840e99afc6cf4cc800869" #cloudflare_list.%[2]s.id
  	comment    = "%[3]s"
  	redirect = {
  		source_url = "www.site1.com/test"
  		target_url = "https://cloudflare.com"
  		status_code = 301
  	}
  }
