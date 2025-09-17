
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[2]s"
  expected_body = "alive"
  expected_codes = "2xx"
}