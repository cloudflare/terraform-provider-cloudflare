
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[2]s"
  expected_body = ""
  expected_codes = "2xx"
  description = "we don't want to check for a given body"
}