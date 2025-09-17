resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "My ruleset %[2]s"
  phase   = "http_request_transform"
  kind    = "zone"
  rules {
    expression = "ip.src eq 1.1.1.1"
    action     = "rewrite"
    action_parameters {
      uri {
        path {
          value = "/new-path"
        }
        query {
          value = "foo=bar"
        }
      }
    }
  }
}