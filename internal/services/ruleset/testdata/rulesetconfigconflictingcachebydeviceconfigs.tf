
  resource "cloudflare_ruleset" "%[1]s" {
	zone_id     = "%[2]s"
    name        = "%[1]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_config_settings"

    rules =[ {
		action = "set_cache_settings"
		action_parameters = {
    cache_key = {
    cache_by_device_type = true
			custom_key = {
    user = {
    geo = false
  }
  }
  }
  }
		expression  = "true"
		description = "do conflicting cache things"
		enabled     = true
	  }]
  }