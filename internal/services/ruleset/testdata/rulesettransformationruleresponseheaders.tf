
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_response_headers_transform"

  rules = [{
    action = "rewrite"
    action_parameters = {
      headers = {
        example1 = {
          operation = "set"
          value     = "my-http-header-value1"
        },
        example2 = {
          operation  = "set"
          expression = "cf.zone.name"
        },
        example3 = {
          operation = "remove"
        }
      }
    }

    expression  = "true"
    description = "example header transformation rule"
    enabled     = false
  }]
}
