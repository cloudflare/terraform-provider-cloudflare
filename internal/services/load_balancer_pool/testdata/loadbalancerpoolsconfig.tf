
resource "cloudflare_load_balancer_pool" "pool1" {
	account_id = "%[2]s"
	name = "pool1"
	origins = [{
		name = "example-1"
		address = "example.com"
		enabled = true
	}]
}

resource "cloudflare_load_balancer_pool" "pool2" {
	account_id = "%[2]s"
	name = "pool2"
	origins = [{
		name = "example-2"
		address = "example.com"
		enabled = true
	}]
}

data "cloudflare_load_balancer_pools" "%[1]s" {
	account_id = "%[2]s"

	depends_on = ["cloudflare_load_balancer_pool.pool1", "cloudflare_load_balancer_pool.pool2"]
}
