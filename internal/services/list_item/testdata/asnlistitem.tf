
  # TODO: once cloudflare_list is working, uncomment the
  # resource
  #
  # resource "cloudflare_list" "%[2]s" {
  #   account_id          = "%[4]s"
  #   name                = "%[2]s"
  #   description         = "list named %[2]s"
  #   kind                = "asn"
  # }

  resource "cloudflare_list_item" "%[1]s" {
    account_id = "%[4]s"
  	list_id    = "c914cfbebb294b15a9ad40113545f23b" #cloudflare_list.%[2]s.id
  	asn = 1
  	comment    = "%[3]s"
  }
