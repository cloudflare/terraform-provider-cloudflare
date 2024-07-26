
resource "cloudflare_ruleset" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_request_late_transform"

  rules = [{
    action = "rewrite"
    action_parameters = {
    headers = [{
        name      = "example1"
        operation = "set"
        value     = "my-http-header-value1"
        },
        {
          name       = "example2"
          operation  = "set"
          expression = "cf.zone.name"
        },
        {
          name      = "example3"
          operation = "remove"
      }]
  }

    expression  = "true"
    description = "example header transformation rule"
    enabled     = false
  }]
}
