
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_log_custom_fields"

  rules = [{
    action = "log_custom_field"
    action_parameters = {
      request_fields = [
        { name = "content-type" },
        { name = "x-forwarded-for" },
        { name = "host" }
      ]
      response_fields = [
        { name = "server" },
        { name = "content-type" },
        { name = "allow" },
      ]
      cookie_fields = [
        { name = "__ga" },
        { name = "accountNumber" },
        { name = "__cfruid" }
      ]
    }

    expression  = "true"
    description = "%[1]s log custom fields rule"
    enabled     = true
  }]
}
