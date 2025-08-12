variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_log_custom_fields"
  kind    = "zone"
  rules = [
    {
      expression = "true"
      action     = "log_custom_field"
      action_parameters = {
        cookie_fields = [
          {
            name = "__cfruid"
          }
        ]
      }
    }
  ]
}
