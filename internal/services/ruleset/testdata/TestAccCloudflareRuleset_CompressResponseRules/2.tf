variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_response_compression"
  kind    = "zone"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "compress_response"
      action_parameters = {
        algorithms = [
          {
            name = "brotli"
          },
          {
            name = "gzip"
          }
        ]
      }
    }
  ]
}