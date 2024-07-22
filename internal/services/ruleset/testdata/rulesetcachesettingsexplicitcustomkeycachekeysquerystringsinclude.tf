
	resource "cloudflare_ruleset" "%[1]s" {
		zone_id     = "%[2]s"
    	name        = "%[1]s"
		description = "set cache settings for the request"
		kind        = "zone"
		phase       = "http_request_cache_settings"
		rules =[ {
		  action      = "set_cache_settings"
		  description = "example"
		  enabled     = true
		  expression  = "(http.host eq \"example.com\" and starts_with(http.request.uri.path, \"/example\"))"
		  action_parameters = {
    cache = true
			edge_ttl = {
    mode    = "override_origin"
			  default = 7200
  }
			cache_key = {
    ignore_query_strings_order = true
			  custom_key = {
    query_string = {
    include = ["another_example"]
  }
  }
  }
  }
		}]
	  }
	