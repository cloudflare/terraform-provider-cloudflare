variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_request_late_transform"
  kind    = "zone"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "rewrite"
      action_parameters = {
        headers = {
          "my-first-header" = {
            operation = "set"
            value     = "my-first-header-value"
          }
          "my-second-header" = {
            operation  = "set"
            expression = "ip.src"
          }
          "my-third-header" = {
            operation = "remove"
          }
        }
      }
    }
  ]
}
