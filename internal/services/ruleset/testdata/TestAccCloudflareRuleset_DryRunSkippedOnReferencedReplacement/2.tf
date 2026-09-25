variable "zone_id" {}

resource "cloudflare_ruleset" "my_custom_ruleset" {
  zone_id = var.zone_id
  name    = "My renamed custom ruleset"
  kind    = "custom"
  phase   = "http_ratelimit"

  rules = [{
    action      = "block"
    expression  = "http.host eq \"example.com\""
    description = "Rate limit example.com"

    ratelimit = {
      characteristics     = ["ip.src", "cf.colo.id"]
      period              = 10
      requests_per_period = 100
      mitigation_timeout  = 10
    }
  }]
}

resource "cloudflare_ruleset" "my_entrypoint_ruleset" {
  zone_id = var.zone_id
  name    = "My entrypoint ruleset"
  kind    = "zone"
  phase   = "http_ratelimit"

  rules = [{
    action      = "execute"
    expression  = "true"
    description = "Execute my custom ruleset"

    action_parameters = {
      id = cloudflare_ruleset.my_custom_ruleset.id
    }
  }]
}
