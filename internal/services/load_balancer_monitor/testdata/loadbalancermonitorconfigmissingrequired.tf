
resource "cloudflare_load_balancer_monitor" "test" {
	account_id = "%[1]s"
	description = "this is a wrong config"
}