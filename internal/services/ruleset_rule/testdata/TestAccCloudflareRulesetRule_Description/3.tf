variable "zone_id" {}

resource "cloudflare_ruleset" "block_external_traffic" {
  kind    = "zone"
  name    = "Block external traffic"
  phase   = "http_request_firewall_custom"
  zone_id = var.zone_id
  lifecycle {
    ignore_changes = [rules]
  }
}

resource "cloudflare_ruleset_rule" "allow_rancher" {
  ruleset_id  = cloudflare_ruleset.block_external_traffic.id
  description = "My rule description"
  action      = "skip"
  action_parameters = {
    ruleset = "current"
  }
  enabled     = true
  expression  = "(starts_with(http.host, \"provisioning\") and ip.src eq 151.251.76.61)"

  zone_id = var.zone_id
  position = {
    index = 1
  }
}
