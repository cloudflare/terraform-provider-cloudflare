
  resource "cloudflare_managed_transforms" "%[1]s" {
	zone_id  = "%[2]s"
	managed_request_headers = [{
		id = "add_true_client_ip_headers"
		enabled = true
	},
    {
    id = "add_visitor_location_headers"
		enabled = true
    }]


	managed_response_headers = [{
		id = "add_security_headers"
		enabled = true
	}]
  }