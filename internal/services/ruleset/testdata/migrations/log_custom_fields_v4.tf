resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "Log Custom Fields Ruleset %[2]s"
  phase   = "http_log_custom_fields"
  kind    = "zone"

  rules {
    expression = "true"
    action     = "log_custom_field"
    description = "Log cookie fields"

    action_parameters {
      cookie_fields = ["session_id", "user_token", "tracking_id"]
    }
  }

  rules {
    expression = "true"
    action     = "log_custom_field"
    description = "Log request fields for API"

    action_parameters {
      request_fields = ["cf.bot_score", "http.user_agent"]
    }
  }

  rules {
    expression = "true"
    action     = "log_custom_field"
    description = "Log response fields for errors"

    action_parameters {
      response_fields = ["cf.ray_id", "cf.colo"]
    }
  }
}
