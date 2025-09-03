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
        raw_response_fields = [
          {
            name = "allow"
          },
          {
            name                = "content-type"
            preserve_duplicates = false
          },
          {
            name                = "server"
            preserve_duplicates = true
          }
        ]
        request_fields = [
          {
            name = "content-type"
          }
        ]
        response_fields = [
          {
            name = "access-control-allow-origin"
          },
          {
            name                = "connection"
            preserve_duplicates = false
          },
          {
            name                = "set-cookie"
            preserve_duplicates = true
          }
        ]
        transformed_request_fields = [
          {
            name = "host"
          }
        ]
      }
    }
  ]
}

data "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  id      = cloudflare_ruleset.my_ruleset.id
}
