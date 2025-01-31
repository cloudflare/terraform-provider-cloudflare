
  resource "cloudflare_magic_wan_static_route" "%[1]s" {
	account_id = "%[3]s"
	prefix = "10.100.0.0/24"
	nexthop = "10.0.0.0"
	priority = "100"
	description = "%[2]s"
	weight = %[4]d
	colo_regions = ["APAC"]
	colo_names = ["den01"]
  }