resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "SBFM Ruleset %[2]s"
  phase   = "http_request_firewall_custom"
  kind    = "zone"

  rules {
    expression = "cf.bot_management.score lt 30"
    action     = "managed_challenge"
    description = "Challenge suspected bots"
  }

  rules {
    expression = "(not cf.bot_management.verified_bot) and (cf.bot_management.score lt 10)"
    action     = "block"
    description = "Block known bad bots"
  }

  rules {
    expression = "(http.request.uri.path contains \"/api\") and (cf.bot_management.score lt 50)"
    action     = "js_challenge"
    description = "JS challenge for API endpoints"
  }
}
