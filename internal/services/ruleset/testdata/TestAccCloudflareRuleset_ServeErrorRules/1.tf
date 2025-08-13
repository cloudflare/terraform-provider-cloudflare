variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_custom_errors"
  kind    = "zone"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "serve_error"
      action_parameters = {
        content      = "1xxx error occurred"
        content_type = "text/plain"
      }
    }
  ]
}
