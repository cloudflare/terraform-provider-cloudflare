variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_response_headers_transform"
  kind    = "zone"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "rewrite"
      action_parameters = {
        headers = {
          "my-first-header" = {
            operation = "add"
            value     = "my-first-header-value"
          }
          "my-second-header" = {
            operation  = "add"
            expression = "http.host"
          }
          "my-third-header" = {
            operation = "set"
            value     = "my-third-header-value"
          }
          "my-fourth-header" = {
            operation  = "set"
            expression = "ip.src"
          }
          "my-fifth-header" = {
            operation = "remove"
          }
        }
      }
    }
  ]
}

data "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  id      = cloudflare_ruleset.my_ruleset.id
}
