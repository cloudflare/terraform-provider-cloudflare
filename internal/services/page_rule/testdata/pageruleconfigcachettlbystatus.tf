
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s"
	actions = [{
		cache_ttl_by_status = [{
			codes = "200-299"
			ttl = 300
		},
    {
    codes = "300-399"
			ttl = 60
    },
    {
    codes = "400-403"
			ttl = -1
    },
    {
    codes = "404"
			ttl = 30
    },
    {
    codes = "405-499"
			ttl = -1
    },
    {
    codes = "500-599"
			ttl = 0
    }]
	}]
}