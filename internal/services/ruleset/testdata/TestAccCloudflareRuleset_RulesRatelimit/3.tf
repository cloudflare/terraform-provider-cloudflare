variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_ratelimit"
  kind    = "zone"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "block"
      ratelimit = {
        characteristics            = ["cf.colo.id", "ip.src"]
        period                     = 60
        counting_expression        = "ip.src eq 2.2.2.2"
        mitigation_timeout         = 600
        requests_to_origin         = true
        score_per_period           = 400
        score_response_header_name = "my-score"
      }
    }
  ]
}

data "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  id      = cloudflare_ruleset.my_ruleset.id
}
