
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_managed"

    rules =[ {
      action = "skip"
      action_parameters = {
    ruleset = "current"
  }
      description = "not this zone"
      expression = "http.host eq \"%[4]s\""
      enabled = true
	  logging = {
    enabled = true
  }
    },
    {
    action = "skip"
      action_parameters = {
    phases = ["http_ratelimit", "http_request_firewall_managed"]
  }
      expression = "http.request.uri.path contains \"/skip-phase/\""
      description = ""
      enabled = true
	  logging = {
    enabled = true
  }
    },
    {
    action = "skip"
      action_parameters = {
    products = ["zoneLockdown", "uaBlock"]
  }
      expression = "http.request.uri.path contains \"/skip-products/\""
      description = ""
      enabled = true
	  logging = {
    enabled = true
  }
    }]


  }