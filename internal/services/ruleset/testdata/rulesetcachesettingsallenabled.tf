
  resource "cloudflare_ruleset" "%[1]s" {
	zone_id     = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_cache_settings"

    rules =[ {
      action = "set_cache_settings"
      action_parameters = {
    additional_cacheable_ports = [8443]
		edge_ttl = {
    mode = "override_origin"
			default = 60
			status_code_ttl =[ {
				status_code = 200
				value = 50
			},
    {
    status_code_range =[ {
					from = 201
					to = 300
				}]
				value = 30
    }]
  }
		browser_ttl = {
    mode = "respect_origin"
  }
		serve_stale = {
    disable_stale_while_updating = true
  }
		respect_strong_etags = true
		read_timeout = 2000
		cache_key = {
    ignore_query_strings_order = false
			cache_deception_armor = true
			custom_key = {
    query_string = {
    exclude = ["*"]
  }
				header = {
    include = ["habc", "hdef"]
					check_presence = ["habc_t", "hdef_t"]
					exclude_origin = true
  }
				cookie = {
    include = ["cabc", "cdef"]
					check_presence = ["cabc_t", "cdef_t"]
  }
				user = {
    device_type = true
					geo = false
  }
				host = {
    resolved = true
  }
  }
  }
		origin_cache_control = true
		origin_error_page_passthru = false
  }
	  expression = "true"
	  description = "%[1]s set cache settings rule"
	  enabled = true
    }]
  }