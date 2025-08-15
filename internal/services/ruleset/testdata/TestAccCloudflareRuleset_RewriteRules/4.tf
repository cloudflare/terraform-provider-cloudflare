variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_request_transform"
  kind    = "zone"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "rewrite"
      action_parameters = {
        uri = {
          path = {
            value = "/foo"
          }
        }
      }
    }
  ]
}
