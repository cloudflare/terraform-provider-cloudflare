
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[2]s"
  name        = "%[1]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_request_dynamic_redirect"

  rules = [{
    action = "redirect"
    action_parameters = {
      from_value = {
        status_code = 301
        target_url = {
          value = "some_host.com"
        }
        preserve_query_string = true
      }
    }
    expression  = "true"
    description = "Apply redirect from value"
    enabled     = true
  }]
}