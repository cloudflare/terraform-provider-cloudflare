
  resource "cloudflare_list" "%[1]s" {
    account_id = "%[4]s"
    name = "%[2]s"
    description = "%[3]s"
    kind = "hostname"

    item =[ {
      value =[ {
        hostname =[ {
		  url_hostname = "*.google.com"
		}]
      }]
      comment = "hostname one"
    },
    {
    value =[ {
		hostname =[ {
		  url_hostname = "manutd.com"
		}]
	  }]
      comment = "hostname two"
    }]

  }