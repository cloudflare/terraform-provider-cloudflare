
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_response_compression"

  rules = [{
    action = "compress_response"
    action_parameters = {
      algorithms = [{
        name = "brotli"
        },
        {
          name = "default"
      }]
    }

    expression  = "true"
    description = "%[1]s compress response rule"
    enabled     = false
  }]
}