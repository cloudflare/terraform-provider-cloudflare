resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "Literal Expression Test %[2]s"
  phase   = "http_request_firewall_custom"
  kind    = "zone"

  rules {
    action = "block"
    expression = <<EOF
    ip.geoip.country eq "CN" and
    http.host eq "example.com"
    EOF
    description = "Literal heredoc test"
    enabled = true
  }
}