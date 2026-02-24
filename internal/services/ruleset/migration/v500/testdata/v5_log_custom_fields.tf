resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "Log Custom Fields Ruleset %[2]s"
  phase   = "http_log_custom_fields"
  kind    = "zone"

  rules = [
    {
      expression  = "true"
      action      = "log_custom_field"
      description = "Log cookie fields"
      action_parameters = {
        cookie_fields = [
          { name = "session_id" },
          { name = "user_token" },
          { name = "tracking_id" }
        ]
      }
    },
    {
      expression  = "true"
      action      = "log_custom_field"
      description = "Log request fields for API"
      action_parameters = {
        request_fields = [
          { name = "cf.bot_score" },
          { name = "http.user_agent" }
        ]
      }
    },
    {
      expression  = "true"
      action      = "log_custom_field"
      description = "Log response fields for errors"
      action_parameters = {
        response_fields = [
          { name = "cf.ray_id" },
          { name = "cf.colo" }
        ]
      }
    }
  ]
}
