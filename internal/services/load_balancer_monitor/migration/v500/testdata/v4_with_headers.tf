resource "cloudflare_load_balancer_monitor" "%s" {
  account_id     = "%s"
  type           = "https"
  method         = "GET"
  path           = "/health"
  expected_codes = "200"

  header {
    header = "Host"
    values = ["%s"]
  }
}
