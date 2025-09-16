locals {
  test_expression = "ip.geoip.country eq \"CN\""
}

resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "Simple Expression Test %[2]s"
  phase   = "http_request_firewall_custom"
  kind    = "zone"

  rules {
    action = "block"
    expression = <<EOF
    ${local.test_expression}
    EOF
    description = "Simple heredoc test"
    enabled = true
  }
}