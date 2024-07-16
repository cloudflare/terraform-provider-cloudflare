
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id     = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_log_custom_fields"

    rules =[ {
      action = "log_custom_field"
      action_parameters = {
    request_fields = [
          "content-type",
          "x-forwarded-for",
          "host"
        ]
        response_fields = [
          "server",
          "content-type",
          "allow"
        ]
        cookie_fields = [
          "__ga",
          "accountNumber",
          "__cfruid"
        ]
  }

      expression = "true"
      description = "%[1]s log custom fields rule"
      enabled = true
    }]
  }