resource "cloudflare_load_balancer_monitor" "%s" {
  account_id     = "%s"
  type           = "http"
  expected_codes = "200"
}
