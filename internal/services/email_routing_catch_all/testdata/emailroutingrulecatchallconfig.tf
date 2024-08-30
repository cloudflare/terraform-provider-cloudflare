resource "cloudflare_email_routing_catch_all" "%[1]s" {
	zone_id = "%[2]s"
	enabled = "%[3]t"
	name = "terraform rule catch all"

	matchers = [{
		type  = "all"
	}]

	actions = [{
		type = "forward"
		value = ["destinationaddress@example.net"]
	}]
}
