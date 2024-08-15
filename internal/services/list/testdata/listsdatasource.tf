
	resource "cloudflare_list" "%[1]s" {
		account_id = "%[2]s"
		name = "%[1]s"
		description = "example list"
		kind = "ip"

		item =[ {
		  value =[ {
			ip = "192.0.2.0"
		  }]
		  comment = "one"
		},
    {
    value =[ {
			ip = "192.0.2.1"
		  }]
		  comment = "two"
    },
    {
    value =[ {
			ip = "192.0.2.2"
		  }]
		  comment = "three"
    },
    {
    value =[ {
			ip = "192.0.2.3"
		  }]
		  comment = "four"
    }]



	  }

data "cloudflare_lists" "%[1]s" {
  account_id = "%[2]s"
  depends_on = [ cloudflare_list.%[1]s ]
}