
  resource "cloudflare_list" "%[1]s" {
    account_id = "%[4]s"
    name = "%[2]s"
    description = "%[3]s"
    kind = "ip"

    item =[ {
		value =[ {
		  ip = "192.0.2.0"
		}]
		comment = "one"
	  }]
  }