variable "zone_id" {}

resource "cloudflare_ruleset" "my_entrypoint_ruleset" {
  zone_id = var.zone_id
  name    = "My entrypoint ruleset"
  kind    = "zone"
  phase   = "http_ratelimit"
}
