
  # TODO: once cloudflare_list is working, uncomment the
  # resource
  #
  # resource "cloudflare_list" "%[2]s" {
  #   account_id          = "%[4]s"
  #   name                = "%[2]s"
  #   description         = "list named %[2]s"
  #   kind                = "redirect"
  # }

  resource "cloudflare_list_item" "%[1]s" {
    account_id = "%[4]s"
  	list_id    = "a2d365108e1840e99afc6cf4cc800869" # cloudflare_list.%[2]s.id
  	comment    = "%[3]s"
  	redirect = {
  		source_url = "example.com/"
  		target_url = "https://example1.com"
  		status_code = 301
  	}
  }
