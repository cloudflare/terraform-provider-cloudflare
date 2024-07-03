
  resource "cloudflare_list" "%[1]s" {
    account_id = "%[4]s"
    name = "%[2]s"
    description = "%[3]s"
    kind = "ip"

    item {
	  value {
	    ip = "192.0.2.2"
	  }
	  comment = "three"
	}

	item {
      value {
        ip = "192.0.2.0"
      }
      comment = "one"
    }

	item {
	  value {
		ip = "192.0.2.3"
	  }
	  comment = "four"
	}

    item {
      value {
        ip = "192.0.2.1"
      }
      comment = "two"
    }
  }