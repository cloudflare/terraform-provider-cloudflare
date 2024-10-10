
  # TODO: once cloudflare_list is working, uncomment the
  # resource
  #
  # resource "cloudflare_list" "%[2]s" {
  #   account_id          = "%[4]s"
  #   name                = "%[2]s"
  #   description         = "list named %[2]s"
  #   kind                = "redirect"
  # }

  resource "cloudflare_list_item" "%[2]s" {
    account_id = "%[4]s"
  	list_id    = "383a2afe7165491ba63733412af6c20e" #cloudflare_list.%[2]s.id
  	ip         = "192.0.2.0"
  	comment    = "%[3]s"
  }
